document.getElementById("logoutBtn").addEventListener("click", function () {
    sessionStorage.removeItem("token"); // Hapus token
    alert("Anda telah logout!");
    window.location.href = "login.html"; // Redirect ke halaman login
  });
  