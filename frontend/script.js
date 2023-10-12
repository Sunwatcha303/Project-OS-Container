// เพิ่มรีวิวของหนัง
document.getElementById("review-form").addEventListener("submit", function (event) {
    event.preventDefault();
    
    const userName = document.getElementById("user-name").value;
    const rating = parseInt(document.getElementById("rating").value);
    const userReview = document.getElementById("user-review").value;

    // ส่งข้อมูลไปยังเซิร์ฟเวอร์เพื่อบันทึกลงในฐานข้อมูล
    // ตรวจสอบความถูกต้องของข้อมูลก่อนส่ง

    // หลังจากบันทึกเสร็จสิ้น, รีเซ็ตคะแนนดาวและแสดงการรีเซ็ต
    resetStars();
    displayReview(userName, rating, userReview);

    // ล้างฟอร์ม
    // document.getElementById("movie-name").value = "";
    document.getElementById("user-name").value = "";
    document.getElementById("rating").value = "";
    document.getElementById("user-review").value = "";
});

// รีเซ็ตคะแนนดาว
function resetStars() {
    const stars = document.querySelectorAll('.star');
    stars.forEach(s => s.classList.remove('active'));
}


// แสดงรีวิวบนหน้าเว็บ
function displayReview(userName, rating, userReview) {
    const reviewDiv = document.createElement("div");
    reviewDiv.className = "review"; // เพิ่มคลาส "review"
    reviewDiv.innerHTML = `ชื่อ ${userName}<p>คะแนน: ${rating} ดาว${userReview} </p>`;
    document.getElementById("movie-list").appendChild(reviewDiv);
}


// เพิ่มเหตุการณ์การคลิกสำหรับดาว
const stars = document.querySelectorAll('.star');
stars.forEach(star => {
    star.addEventListener('click', () => {
        const rating = parseInt(star.getAttribute('data-rating'));

        // ให้ดาวทั้งหมดเปลี่ยนสี
        stars.forEach(s => s.classList.remove('active'));

        // เลือกดาวที่ผู้ใช้คลิกและดาวที่มีคะแนนน้อยกว่าหรือเท่ากับคะแนนที่ผู้ใช้เลือก
        stars.forEach((s, index) => {
            if (index < rating) {
                s.classList.add('active');
            }
        });

        document.getElementById('rating').value = rating;
    });
});


// เพิ่มรีวิวของหนัง
// document.getElementById("review-form").addEventListener("submit", function (event) {
//     event.preventDefault();

//     // คำสั่ง AJAX เพื่อดึงคะแนนเฉลี่ยจากไฟล์ PHP
//     var xhr = new XMLHttpRequest();
//     xhr.open("GET", "getAverageRating.php", true);
//     xhr.onreadystatechange = function () {
//         if (xhr.readyState === 4 && xhr.status === 200) {
//             var averageRating = xhr.responseText;
//             document.getElementById("average").textContent = averageRating;
//         }
//     };
//     xhr.send();
    
//     // เพิ่มรีวิว
//     // ...

//     // ล้างฟอร์ม
//     // ...
// });






