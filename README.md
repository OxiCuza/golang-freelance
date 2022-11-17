<!-- PROJECT SHIELDS -->
[![LinkedIn][linkedin-shield]][linkedin-url]


<!-- PROJECT LOGO -->
<br />
<div align="center" id="readme-top">
  <a href="https://go.dev/">
    <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_LightBlue.png" alt="Logo" width="200" height="200">
  </a>
</div>



<!-- ABOUT THE PROJECT -->
## About The Project

Poin utama mengenai pattern yang digunakan, yaitu mengadopsi repository pattern
dengan memisahkan logika aplikasi dari logika database, hal ini berdampak dalam
kemudahan developer untuk menambahkan fungsionalitas baru ke dalam aplikasi.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* Go
* MySQL / MariaDB
* HttpRouter
* Google UUID
* GoDotEnv
* Go Playground Validator
* Redis

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Requirement

* package [Go](https://go.dev/learn/)
* database MySQL / MariaDB
* import schema (doc/db.sql) to your DBMS

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/OxiCuza/golang-freelance.git
   ```
2. Install go mod (optional)
   ```sh
   go mod tidy
   ```
3. Copy .env.example to .env
   ```sh
   cp .env.example .env
   ```
4. Make sure to set up DB env and REDIS env correctly
5. Example env value for API KEY
   ```sh
   X-API-Key=31UOY+zPB4qMFNBbWntF9I75+dl43RvsUmONnwy7C80=
   ```
6. Run App
   ```sh
   go run main.go
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- API Documentation -->
## API Documentation

https://app.swaggerhub.com/apis/OxiCuza/simple-blog-post/1.0.0

example usage :
* Hit endpoint /users with method GET
   ```shell
   curl -X 'GET' \
     'localhost://3000/api/v1/users' \
     -H 'accept: application/json' \
     -H 'X-API-Key: 31UOY+zPB4qMFNBbWntF9I75+dl43RvsUmONnwy7C80='
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- API Documentation -->
## ERD
https://dbdiagram.io/d/6373898ac9abfc611172cc74

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Oxicusa Gugi Housman - [Telegram](https://t.me/OxiCuza) - email : oxicusa@gmail.com

<p>
    Jika berkenan, mohon beri masukan terhadap source code saya melalui email, sebagai bahan evaluasi kedepannya. Terima Kasih
</p>

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/oxicusa/