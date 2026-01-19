async function searchCustomerByName() {
  const name = document.getElementById("searchName").value;

  const res = await fetch("/customers");
  const customers = await res.json();

  const found = customers.filter(c =>
    c.name.toLowerCase().includes(name.toLowerCase())
  );

  let html = "";
  found.forEach(c => {
    html += `<div>
      <b>${c.name}</b> - ${c.email}
      <button onclick="loadAccounts(${c.id})">Hesapları</button>
    </div>`;
  });

  document.getElementById("result").innerHTML = html;
}

async function loadAccounts(customerId) {
  const res = await fetch(`/customers/${customerId}/accounts`);
  const accounts = await res.json();

  let html = "";
  accounts.forEach(a => {
    html += `
      <div>
        Hesap ID: ${a.ID} | Bakiye: ${a.Balance}
        <button onclick="deposit(${a.ID})">+</button>
        <button onclick="withdraw(${a.ID})">-</button>
      </div>
    `;
  });

  document.getElementById("accounts").innerHTML = html;
}

async function deposit(id) {
  const amount = prompt("Yatırılacak tutar:");
  await fetch(`/accounts/${id}/deposit`, {
    method: "POST",
    headers: {"Content-Type": "application/json"},
    body: JSON.stringify({ amount: Number(amount) })
  });
  alert("Para yatırıldı");
}

async function withdraw(id) {
  const amount = prompt("Çekilecek tutar:");
  await fetch(`/accounts/${id}/withdraw`, {
    method: "POST",
    headers: {"Content-Type": "application/json"},
    body: JSON.stringify({ amount: Number(amount) })
  });
  alert("Para çekildi");
}
