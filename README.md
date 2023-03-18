# Go Web Server

This is a simple Go server that serves static files from a `public` directory. The server can serve HTML, CSS, JS, and image files.

## Installation

1. Clone this repository:
`git clone https://github.com/NaviRocker/go-web-server.git`

2. Change to the directory of the cloned repository:
`cd go-web-server`

3. Start the server:
`go run server.go`


4. Navigate to `http://localhost:8080` in your web browser to view the index.html file.

## Folder Structure
go-web-server/  
├── server.go
└── public/
    ├── index.html
    ├── style.css
    ├── script.js   
    └── img/
        └── example.jpg

- `server.go`: The main Go file that sets up the server.
- `public/`: A directory that contains all the static files served by the server.
- `public/index.html`: The main HTML file that is served as the homepage.
- `public/style.css`: The CSS file that is linked to the `index.html` file.
- `public/script.js`: The JavaScript file that is linked to the `index.html` file.
- `public/img/`: A directory that contains all the image files served by the server.
- `public/img/example.jpg`: An example image file.

## License

This code is licensed under the MIT License. See the [LICENSE](https://github.com/NaviRocker/Go-Web-Server/blob/main/LICENSE) file for details.





