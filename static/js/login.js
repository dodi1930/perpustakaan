document.getElementById("loginForm").addEventListener("submit", function (event) {
    event.preventDefault(); // Mencegah form reload halaman
  
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
  
    fetch("http://localhost:9090/admin/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password })
    })
      .then(response => response.json())
      .then(data => {
        if (data.token) {
          sessionStorage.setItem("token", data.token); // Simpan token ke sessionStorage
          alert("Login berhasil!");
          window.location.href = "admin_dashboard.html"; // Redirect ke halaman admin
        } else {
          alert("Login gagal: " + data.error);
        }
      })
      .catch(error => console.error("Error:", error));
  });
  