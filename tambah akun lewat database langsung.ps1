tambah akun lewat database langsung
INSERT INTO users (name, email, password, role, created_at) 
VALUES ('Admin Baru', 'dodi@cgpt.com', '$2a$12$lFBEPlg.pjn3NllHN1TpLuxF08pbYZR9LBCrM6S079pMjAEhVBRUy', 'admin', NOW());


log in admin
Invoke-RestMethod -Uri "http://localhost:9090/login" -Method Post -ContentType "application/json" -Body '{"email":"dodi@cgpt.com","password":"admin123"}'

Invoke-RestMethod -Uri "http://localhost:9090/login" -Method Post -ContentType "application/json" -Body '{"email":"dodi@cgpt.com","password":"admin123"}'



Invoke-RestMethod -Uri "http://localhost:9090/admin/add-user" -Method Post -Headers @{"Admin-Email"="dodi@cgpt.com"} -ContentType "application/json" -Body '{"name":"User Baru","email":"baru@cgpt.com","password":"password","role":"siswa"}'

edit akun
Invoke-RestMethod -Uri "http://localhost:9090/admin/edit-user/1" -Method Put -Headers @{"Admin-Email"="dodi@cgpt.com"} -ContentType "application/json" -Body '{"name":"Dodi","email":"baru@cgpt.com","role":"siswa"}'

Invoke-RestMethod -Uri "http://localhost:9090/admin/edit-user/10" -Method Put -Headers @{"Admin-Email"="dodi@cgpt.com"} -ContentType "application/json" -Body '{"name":"Dodi Prasetyo","email":"dodi@cgpt","role":"admin"}'


hapus akun
Invoke-RestMethod -Uri "http://localhost:9090/admin/delete-user/10" -Method Delete -Headers @{"Admin-Email"="dodi@cgpt.com"}

Import file
Invoke-RestMethod -Uri "http://localhost:9090/admin/import-users" -Method Post -Headers @{"Admin-Email"="dodi@cgpt.com"} -ContentType "multipart/form-data" -InFile "C:\Users\Dodi\Downloads\template_users.xlsx"
role::text = ANY (ARRAY['siswa'::character varying, 'guru'::character varying, 'admin'::character varying, 'kepala_sekolah'::character varying, 'pustakawan'::character varying]::text[])
