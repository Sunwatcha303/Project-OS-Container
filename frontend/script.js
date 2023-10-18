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





