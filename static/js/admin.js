const token = sessionStorage.getItem("token");

if (!token) {
  alert("Anda harus login terlebih dahulu!");
  window.location.href = "login.html"; // Redirect ke halaman login jika tidak ada token
} else {
  fetch("http://localhost:9090/admin/dashboard", { // GANTI PORT KE 9090
    method: "GET",
    headers: { "Authorization": `Bearer ${token}` }
  })
    .then(response => response.json())
    .then(data => {
      console.log("Data Admin:", data);
      document.getElementById("adminName").innerText = data.admin_name;
    })
    .catch(error => console.error("Error:", error));
}
