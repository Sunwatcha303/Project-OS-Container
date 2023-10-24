// JavaScript to add padding to the content when the header becomes fixed
const header = document.querySelector('header');
const content = document.querySelector('main');

content.style.paddingTop = header.clientHeight + 'px';

// Update padding on window resize
window.addEventListener('resize', () => {
    content.style.paddingTop = header.clientHeight + 'px';
});

// รีเซ็ตคะแนนดาว
function resetStars() {
    const stars = document.querySelectorAll('.star');
    stars.forEach(s => s.classList.remove('active'));
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

//part1 by sun
//ใส่รูป-ชื่อ-คะแนนเฉลี่ย-หมวดหมู่ ของหนัง
async function getMovie() {
    const image = document.getElementById("movie-image");
    const name_movie = document.getElementById("movie-name")
    const score_movie = document.getElementById("average")
    const category_movie = document.getElementById("category")
    //get param ดึงค่า parameter ชื่อ 'id_movie' จาก URL ของหน้าเว็บปัจจุบัน
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id_movie');
    try {
        //ส่งคำขอ GET (การดึงข้อมูล) ไปยัง URL 
        const header = {
            'Api-Key': '1234567890'
        }
        const response = await fetch("http://localhost:8080/project-os-container/movies/" + id, { method: "GET", headers: header });

        if (response.status === 200) {
            //แปลงข้อมูลที่ได้รับจาก response เป็น JSON format โดยใช้ await เพื่อรอให้การแปลงเสร็จสิ้น ซึ่งเป็นการใช้งาน asynchronous.
            const movieData = await response.json();

            //รับข้อมูลบิตแปลงข้อมูลบิตเป็นรูปภาพ
            const base64String = movieData.image_movie;
            const binaryString = atob(base64String);
            const uint8Array = new Uint8Array(binaryString.length);
            for (let i = 0; i < binaryString.length; i++) {
                uint8Array[i] = binaryString.charCodeAt(i);
            }
            const blob = new Blob([uint8Array], { type: "image/jpeg" });
            const urlCreator = window.URL || window.webkitURL;
            image.src = urlCreator.createObjectURL(blob);


            //ชื่อ-คะแนนเฉลี่ย-หมวดหมู่
            name_movie.innerHTML = movieData.movie_name
            score_movie.innerHTML = movieData.movie_score
            category_movie.innerHTML += movieData.category
        } else {
            alert("Error fetching the image.");
        }
    } catch (error) {
        console.error("Error:", error);
    }
}
getMovie();


//part2 by ManW
document.getElementById("review-form").addEventListener("submit", async function (event) {
    event.preventDefault();

    // ส่งข้อมูลไปยังเซิร์ฟเวอร์เพื่อบันทึกลงในฐานข้อมูล
    // ตรวจสอบความถูกต้องของข้อมูลก่อนส่ง

    const userName = document.getElementById("user-name").value;
    const rating = parseInt(document.getElementById("rating").value);
    const userReview = document.getElementById("user-review").value;

    //get param
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id_movie');
    const data = {
        name: userName,
        id_movie: parseInt(id),
        comment: userReview,
        score: rating
    };
    try {
        const header = {
            'Api-Key': '1234567890'
        }
        const response = await fetch("http://localhost:8080/project-os-container/reviews/add", { method: "POST", headers: header, body: JSON.stringify(data) });
        if (response.status === 201) {
            if (document.getElementById("movie-list").innerHTML == "No Comment Review") {
                document.getElementById("movie-list").innerHTML = ""
            }
            displayReview(userName, rating, userReview, new Date())

        } else {
            alert("Error");
        }
    } catch (error) {
        console.error("Error:", error);
    }

    // หลังจากบันทึกเสร็จสิ้น, รีเซ็ตคะแนนดาวและแสดงการรีเซ็ต
    resetStars();
    //displayReview(userName, rating, userReview);

    // ล้างฟอร์ม
    // document.getElementById("movie-name").value = "";
    document.getElementById("user-name").value = "";
    document.getElementById("rating").value = "0";
    document.getElementById("user-review").value = "";

    // location.reload();
});

//part 3 by aum and dear
async function getReviews() {
    //get param
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id_movie');
    const header = {
        'Api-Key': '1234567890'
    }
    fetch("http://localhost:8080/project-os-container/reviews/" + id, { method: "GET", headers: header })
        .then(response => {
            if (response.status === 204) {
                return;
            }
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            //ส่งข้อมูลทั้งหมดไปเข้าฟังก์ชัน displayReview
            if (data == null) {
                document.getElementById("movie-list").innerHTML = "No Comment Review"
                return;
            }
            data.forEach(review => {
                displayReview(review.name, review.score, review.comment, review.create_at);
            });
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });
}
getReviews();

function displayReview(userName, rating, userReview, Time) {
    const now = new Date();
    const reviewTime = new Date(Time);
    const timeDifference = now - reviewTime;

    let formattedDateTime;

    if (timeDifference <= 30000) { // 30 second
        formattedDateTime = "now";
    } else if (timeDifference <= 86400000) {// 1 วันมี 86400000 มิลลิวินาที
        formattedDateTime = reviewTime.toLocaleTimeString('en-US', { timeStyle: 'short' });
    } else {
        formattedDateTime = reviewTime.toLocaleDateString('en-US', { dateStyle: 'long' });
    }

    const reviewDiv = document.createElement("div");
    reviewDiv.className = "review"; // เพิ่มคลาส "review"
    reviewDiv.innerHTML = `Name ${userName} (${rating} star) : <small style="color: #888">${formattedDateTime}</small> </p> <p> ${userReview} </p> `;

    document.getElementById("movie-list").appendChild(reviewDiv);
    getScoreMovie()
}

async function getScoreMovie(){
    const score_movie = document.getElementById("average")
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id_movie');
    try{
        const header = {
            'Api-Key': '1234567890'
        }
        const response = await fetch("http://localhost:8080/project-os-container/movies/score/" + id, {method: "GET",headers: header})
        if (response.status === 200){
            movieData = await response.json()
            score_movie.innerHTML = movieData.movie_score
        }
    }catch(error){
        console.error("Error:", error);
    }
}
