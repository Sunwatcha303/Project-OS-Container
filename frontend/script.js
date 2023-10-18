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

//part1
const image = document.getElementById("movie-image");
const name_movie = document.getElementById("movie-name")
const score_movie = document.getElementById("average")

async function getImage() {
    try {
        const header = {
            'Content-Type': 'application/json',
            'Api_Key': '1234567890'
        }
        const response = await fetch("http://localhost:8080/project-os-container/movie/1", { method: "GET", headers: header }); // Replace '1' with the actual image ID

        if (response.status === 200) {
            const movieData = await response.json();

            const base64String = movieData.image_movie;
            const binaryString = atob(base64String);
            const uint8Array = new Uint8Array(binaryString.length);
            for (let i = 0; i < binaryString.length; i++) {
                uint8Array[i] = binaryString.charCodeAt(i);
            }

            const blob = new Blob([uint8Array], { type: "image/jpeg" }); // Adjust the content type as needed
            const urlCreator = window.URL || window.webkitURL;

            image.src = urlCreator.createObjectURL(blob);

            console.log(image.src)
            name_movie.innerHTML = movieData.movie_name
            score_movie.innerHTML = movieData.movie_score
        } else {
            alert("Error fetching the image.");
        }
    } catch (error) {
        console.error("Error:", error);
    }
}

getImage();