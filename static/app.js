// static/app.js

// Küçük yardımcı: farklı JSON alan adlarını normalize edelim
function pick(obj, keys, fallback = "") {
  for (const k of keys) {
    if (obj && obj[k] !== undefined && obj[k] !== null) return obj[k];
  }
  return fallback;
}

function toLowerSafe(x) {
  return (x ?? "").toString().toLowerCase();
}

async function safeJson(res) {
  try { return await res.json(); } catch { return null; }
}

// GLOBAL STATE (seçili müşteri)
let selectedCustomerId = null;
let selectedCustomerName = "";

// HTML’de bazen result bazen searchResult kullanılmış olabiliyor
function getSearchResultEl() {
  return document.getElementById("searchResult") || document.getElementById("result");
}

// 1) Müşteri Arama (isimle)
async function searchCustomerByName() {
  const q = document.getElementById("searchName").value.trim();
  const resultEl = getSearchResultEl();
  const accountsEl = document.getElementById("accounts");

  if (!resultEl) {
    console.error("Arama sonucu alanı bulunamadı. (searchResult veya result id'si yok)");
    return;
  }

  resultEl.innerHTML = "";
  accountsEl.innerHTML = "";

  if (!q) {
    resultEl.innerHTML = `<div class="text-muted">Aramak için bir isim yaz.</div>`;
    return;
  }

  // Tercih: backend search endpoint
  let customers = [];
  let res = await fetch(`/customers/search?q=${encodeURIComponent(q)}`);

  if (res.ok) {
    customers = (await safeJson(res)) || [];
  } else {
    // Fallback: tüm müşterileri çekip filtrele
    res = await fetch("/customers");
    const all = (await safeJson(res)) || [];
    customers = all.filter(c => toLowerSafe(pick(c, ["Name","name"])) .includes(toLowerSafe(q)));
  }

  if (!Array.isArray(customers) || customers.length === 0) {
    resultEl.innerHTML = `<div class="alert alert-warning mb-0">Sonuç bulunamadı.</div>`;
    return;
  }

  // Kart gibi göster
  let html = "";
  customers.forEach(c => {
    const id = pick(c, ["ID", "id"]);
    const name = pick(c, ["Name", "name"]);
    const email = pick(c, ["Email", "email"]);

    html += `
      <div class="p-3 mb-2 border rounded bg-white d-flex align-items-center justify-content-between">
        <div>
          <div class="fw-bold">${name}</div>
          <div class="text-muted small">${email} (ID: ${id})</div>
        </div>
        <div class="d-flex gap-2">
          <button class="btn btn-sm btn-outline-primary"
            onclick="selectCustomerAndLoad(${id}, '${String(name).replace(/'/g, "\\'")}')">
            Hesabı Göster
          </button>
          <button class="btn btn-sm btn-outline-danger" onclick="deleteCustomer(${id})">Sil</button>
        </div>
      </div>
    `;
  });

  resultEl.innerHTML = html;
}

// Seçili müşteriyi kaydet + hesaplarını çek
async function selectCustomerAndLoad(customerId, customerName) {
  selectedCustomerId = Number(customerId);
  selectedCustomerName = customerName || "";
  await loadAccounts(selectedCustomerId, selectedCustomerName);
}

// 2) Müşteri silme (onaylı)
async function deleteCustomer(id) {
  const ok = confirm(`ID ${id} olan müşteriyi silmek istediğine emin misin?`);
  if (!ok) return;

  const res = await fetch(`/customers/${id}`, { method: "DELETE" });
  if (!res.ok) {
    const data = await safeJson(res);
    alert(data?.error || "Silme başarısız");
    return;
  }
  alert("Müşteri silindi ✅");

  // Eğer seçili müşteri silindiyse ekranı temizle
  if (selectedCustomerId === Number(id)) {
    selectedCustomerId = null;
    selectedCustomerName = "";
    const accountsEl = document.getElementById("accounts");
    if (accountsEl) accountsEl.innerHTML = "";
  }

  // arama listesini yenile
  await searchCustomerByName();
}

// 3) Hesapları Listele
// - Her hesabın detayını /accounts/:id/details ile çek
async function loadAccounts(customerId, customerName = "") {
  const accountsEl = document.getElementById("accounts");
  accountsEl.innerHTML = `<div class="text-muted">Hesaplar yükleniyor...</div>`;

  const res = await fetch(`/customers/${customerId}/accounts`);
  const accounts = (await safeJson(res)) || [];

  if (!res.ok) {
    accountsEl.innerHTML = `<div class="text-danger">Hesaplar getirilemedi.</div>`;
    return;
  }

  if (!Array.isArray(accounts) || accounts.length === 0) {
    accountsEl.innerHTML = `<div class="alert alert-info mb-0">Bu müşterinin hesabı yok.</div>`;
    return;
  }

  // Her hesap için details çekelim (isim + lastAction + balance)
  const detailList = [];
  for (const a of accounts) {
    const accId = pick(a, ["ID", "id"]);
    const dres = await fetch(`/accounts/${accId}/details`);
    const d = await safeJson(dres);
    if (dres.ok && d) detailList.push(d);
  }

  // UI
  let html = `<div class="mb-2 fw-bold">Seçili Müşteri: ${customerName || ("Customer ID: " + customerId)}</div>`;

  detailList.forEach(d => {
    html += renderAccountCard(d);
  });

  accountsEl.innerHTML = html;
}

function renderAccountCard(d) {
  // d: {customerName, accountId, customerId, balance, lastAction}
  const customerName = pick(d, ["customerName"]);
  const accountId = pick(d, ["accountId"]);
  const customerId = pick(d, ["customerId"]);
  const balance = pick(d, ["balance"], 0);
  const lastAction = pick(d, ["lastAction"], "Henüz işlem yok");

  return `
    <div class="bg-dark text-white p-3 rounded mb-3">
      <div class="fw-bold fs-5 mb-2">${customerName}</div>

      <div><b>Account ID:</b> ${accountId}</div>
      <div><b>Customer ID:</b> ${customerId}</div>
      <div><b>Bakiye:</b> ${balance} ₺</div>

      <div class="mt-2"><b>Son İşlem:</b> ${lastAction}</div>

      <div class="mt-3 d-flex gap-2">
        <button class="btn btn-sm btn-success" onclick="deposit(${accountId})">Para Yatır</button>
        <button class="btn btn-sm btn-warning" onclick="withdraw(${accountId})">Para Çek</button>
      </div>
    </div>
  `;
}

// 4) Para Yatır / Çek
// - işlem sonrası hesabı yenile (seçili müşteriye göre)
async function deposit(accountId) {
  const amountStr = prompt("Yatırılacak tutar:");
  if (amountStr === null) return;

  const amount = Number(amountStr);
  if (!Number.isFinite(amount) || amount <= 0) {
    alert("Geçersiz tutar");
    return;
  }

  const res = await fetch(`/accounts/${accountId}/deposit`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ amount })
  });

  const data = await safeJson(res);
  if (!res.ok) {
    alert(data?.error || "Para yatırma başarısız");
    return;
  }

  alert("Para yatırıldı ✅");

  // seçili müşteri varsa onun hesaplarını tekrar çek
  if (selectedCustomerId) {
    await loadAccounts(selectedCustomerId, selectedCustomerName);
  }
}

async function withdraw(accountId) {
  const amountStr = prompt("Çekilecek tutar:");
  if (amountStr === null) return;

  const amount = Number(amountStr);
  if (!Number.isFinite(amount) || amount <= 0) {
    alert("Geçersiz tutar");
    return;
  }

  const res = await fetch(`/accounts/${accountId}/withdraw`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ amount })
  });

  const data = await safeJson(res);
  if (!res.ok) {
    alert(data?.error || "Para çekme başarısız");
    return;
  }

  alert("Para çekildi ✅");

  if (selectedCustomerId) {
    await loadAccounts(selectedCustomerId, selectedCustomerName);
  }
}

  const searchCustomer = searchCustomerByName;