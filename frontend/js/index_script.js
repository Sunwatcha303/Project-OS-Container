async function getListMovies() {
    try {
        const header = {
            'Api-Key': '1234567890'
        }
        const response = await fetch("http://localhost:8080/project-os-container/movies/all", { method: "GET", headers: header });

        if (response.status === 200) {
            const listMovies = await response.json();
            for (const movie of listMovies) {
                //create new tag
                const movieDiv = document.createElement("div");
                const reviewHref = document.createElement("a");
                const childMovieDiv = document.createElement("div");
                const imageMovie = document.createElement("img");

                //link to review
                reviewHref.href = "/frontend/review.html?id_movie=" + movie.id_movie;

                imageMovie.id = "movie-image";
                imageMovie.alt = "Image";
                imageMovie.height = 300;

                const base64String = movie.image_movie;
                const binaryString = atob(base64String);
                const uint8Array = new Uint8Array(binaryString.length);
                for (let i = 0; i < binaryString.length; i++) {
                    uint8Array[i] = binaryString.charCodeAt(i);
                }
                const blob = new Blob([uint8Array], { type: "image/jpeg" });
                const urlCreator = window.URL || window.webkitURL;

                imageMovie.src = urlCreator.createObjectURL(blob);
                console.log(imageMovie.src)
                childMovieDiv.innerHTML = movie.movie_name;

                //append to tage
                reviewHref.appendChild(imageMovie);

                movieDiv.appendChild(reviewHref);
                movieDiv.appendChild(childMovieDiv);
                document.getElementById("movies-list").appendChild(movieDiv);
            }

        } else {
            alert("Error.");
        }
    } catch (error) {
        console.error("Error:", error);
    }
}
getListMovies();