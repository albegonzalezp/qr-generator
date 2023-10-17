# QR Code Generator

Create QR codes effortlessly with our easy-to-use QR code generator. Whether you want to render a QR code, download it, or send it via SMTP, this tool has you covered.

## Getting Started
To get started, follow these simple steps:

1. **Clone the Repository**: 
   ```sh
   git clone https://github.com/zetacoder/qr-generator.git
   cd qr-generator/backend

2. **Install Dependencies**:
   ```sh
   go mod tidy

3. **Configure the directories and env variables**:
- Go to **/qr-generator/backend**
- Check `.env.example` and set the email configuration you want to use and directories to load templates.
- IMPORTANT: To be able to send emails trough SMTP, you must configure the password in your account. Click [here](https://support.google.com/mail/answer/185833?hl=en#app-passwords) to know how to do it. After that add it to **PASSWORD** in `.env.example`.
- Once all done, save the file as `.env` so can be readed by the app.


5.` **Run the Application**:
- Open the project in your preferred code editor.
- Go to /backend directory and run: `go run .`
- This will serve all the frontend elements from the backend. Server-side rendering.


6. **Access the Web Interface**:
Open a web browser and navigate to http://localhost:8080 (by default if running in local machine).`


7. **Fill out the Form:**
- Complete the form with the necessary information.
- Choose to render, download, or send the QR code via SMTP.


8. **DONE!**
- Use the QR for the purpose that fit you best!
--------------------------------------------------

## **Features**
1. Easy-to-Use: A user-friendly interface for generating QR codes.
2. Render QR Codes: View the QR code directly in your browser.
3. Download: Save the QR code as an image file for offline use.
4. Email Sending: Send the QR code via email using SMTP.

## **Contributing**
If you'd like to contribute to this project, please follow these guidelines:

1. Fork the repository
2. Create a new branch for your feature or bug fix
3. Commit your changes and push them to your fork
4. Submit a pull request

## **Contact**
If you have any questions or feedback, please contact Alberto Gonzalez at albertoigp93@gmail.com
Enjoy creating QR codes with our simple and efficient QR code generator!
