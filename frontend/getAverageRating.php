<?php
// เชื่อมต่อกับฐานข้อมูล MySQL
$servername = "localhost";  // หรือชื่อเซิร์ฟเวอร์ของคุณ
$username = "ชื่อผู้ใช้";  // ชื่อผู้ใช้ฐานข้อมูล
$password = "รหัสผ่าน";  // รหัสผ่านฐานข้อมูล
$database = "ชื่อฐานข้อมูล";  // ชื่อฐานข้อมูล

$conn = new mysqli($servername, $username, $password, $database);

// ตรวจสอบการเชื่อมต่อ
if ($conn->connect_error) {
    die("การเชื่อมต่อล้มเหลว: " . $conn->connect_error);
}

// คำสั่ง SQL เพื่อดึงคะแนนเฉลี่ย
$sql = "SELECT AVG(rating) AS average_rating FROM reviews";
$result = $conn->query($sql);

if ($result->num_rows > 0) {
    $row = $result->fetch_assoc();
    $averageRating = $row["average_rating"];
    echo $averageRating;
} else {
    echo "ไม่พบข้อมูลคะแนนเฉลี่ย";
}

// ปิดการเชื่อมต่อกับฐานข้อมูล
$conn->close();
?>
