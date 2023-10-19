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
    document.getElementById("rating").value = "0";
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
    reviewDiv.innerHTML = `Name ${userName} (${rating} star)</p><p> ${userReview} </p>`;
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

//aum
fetch('http://localhost:8080/project-os-container/')
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        displayReview(data.code, data.status, "kuy savelnwza");
        data.forEach(review => {
            displayReview(review.code, review.status, review.message);
        });
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
    });


// Define the API endpoint URL
// const url = 'http://localhost:8080/project-os-container/'; // Replace with your API endpoint URL

// // Define the JSON data you want to send
// const data = {
//     key1: 'value1',
//     key2: 'value2'
// };

// // Convert the data to a JSON string
// const jsonData = JSON.stringify(data);

// // Set the headers to indicate that you're sending JSON data
// const headers = {
//     'Content-Type': 'application/json'
// };

// // Create a POST request using the fetch API
// fetch(url, {
//     method: 'POST',
//     headers: headers,
//     body: jsonData
// })
//     .then(response => {
//         // Check for a successful response (status code 2xx indicates success)
//         if (response.ok) {
//             return response.json(); // Parse the response as JSON
//         } else {
//             throw new Error('Request failed with status ' + response.status);
//         }
//     })
//     .then(responseData => {
//         console.log('Request was successful.');
//         console.log('Response data:', responseData);
//     })
//     .catch(error => {
//         console.error('Request error:', error);
//     });

//part1 by sun
const image = document.getElementById("movie-image");
const name_movie = document.getElementById("movie-name")
const score_movie = document.getElementById("average")
const category_movie = document.getElementById("category")

async function getImage() {
    try {
        const header = {
            'Api_Key': '1234567890'
        }
        const response = await fetch("http://localhost:8080/project-os-container/movie/1", { method: "GET", headers: header });

        if (response.status === 200) {
            const movieData = await response.json();

            const base64String = movieData.image_movie;
            const binaryString = atob(base64String);
            const uint8Array = new Uint8Array(binaryString.length);
            for (let i = 0; i < binaryString.length; i++) {
                uint8Array[i] = binaryString.charCodeAt(i);
            }

            const blob = new Blob([uint8Array], { type: "image/jpeg" });
            const urlCreator = window.URL || window.webkitURL;

            image.src = urlCreator.createObjectURL(blob);

            name_movie.innerHTML = movieData.movie_name
            score_movie.innerHTML = movieData.movie_score
            category_movie.innerHTML = movieData.category
        } else {
            alert("Error fetching the image.");
        }
    } catch (error) {
        console.error("Error:", error);
    }
}

getImage();





//part2 by Man

document.getElementById("review-form").addEventListener("submit", function (event) {
    event.preventDefault();
    
    // ส่งข้อมูลไปยังเซิร์ฟเวอร์เพื่อบันทึกลงในฐานข้อมูล
    // ตรวจสอบความถูกต้องของข้อมูลก่อนส่ง

    const userName = document.getElementById("user-name").value;
    const rating = parseInt(document.getElementById("rating").value);
    const userReview = document.getElementById("user-review").value; 
    const data = {
        name: userName,
        id_movie:'1',
        comment:userReview,
        score:rating
    };
    try {
        const header = {
            'Api_Key': '1234567890'
        }
        const response = fetch("http://localhost:8080/project-os-container/addReview/1", { method: "POST", headers: header ,body:JSON.stringify(data)});
        if (response.status === 201) {
            console.log("Success")
        } else {
            alert("Error");
        }
    } catch (error) {
        console.error("Error:", error);
    }
    
    // หลังจากบันทึกเสร็จสิ้น, รีเซ็ตคะแนนดาวและแสดงการรีเซ็ต
    resetStars();
    displayReview(userName, rating, userReview);
    
    // ล้างฟอร์ม
    // document.getElementById("movie-name").value = "";
    document.getElementById("user-name").value = "";
    document.getElementById("rating").value = "0";
    document.getElementById("user-review").value = "";
});