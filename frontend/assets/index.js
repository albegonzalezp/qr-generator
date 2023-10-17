
function getQR() {
    if (validateForm() !== true) {
        return;
    }

    // post request to backend: /get-qr using fetch API
    fetch("/get-qr", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            qr_name: document.getElementById("qr-name").value,
            email: document.getElementById("email").value,
            link: document.getElementById("qr-url").value,
            image_type: document.getElementById("qr-extension").value
        })
    })
    .then(response => response.json())
    .then(data => {
        if (data.status !== "success") {
            throw new Error("Something went wrong.");
        }
        // get the image
        let image = document.getElementById("qr")
        // render the qr
        image.src = data.image_path;
        image.name = data.qr_name;
        image.dataset.link = data.link;
        image.dataset.email = data.email;

        // clear the form and input fields
        document.getElementById("qr-input-form").reset();

        // remove the disabled attribute from the buttons id download and send_email 
        document.getElementById("download").removeAttribute("disabled");
        document.getElementById("send_email").removeAttribute("disabled");
    })
    .catch((error) => {
        console.error("Error:", error);
    });
}

function sendQrToEmail() {
    let qrImg = document.getElementById("qr");

    if (qrImg.src === "") {
        alert("Please generate a QR code first.");
        return;
    }

    if (qrImg.dataset.email === "") {
        alert("Please enter a valid email address.");
        return;
    }

    // post request to backend: /send-qr using fetch API
    fetch("/send-qr", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            email: qrImg.dataset.email,
            qr_name: qrImg.name,
            link: qrImg.dataset.link,
            image_path: qrImg.src
        })		
    })
    .then(response => response.json())
    .then(data => {
        if (data.status !== "success") {
            throw new Error("Something went wrong sending the email.");
        }
        // clear the form and input fields
        document.getElementById("qr-input-form").reset();

        alert("Email sent successfully!");
    })
    .catch((error) => {
        console.error("Error:", error);
    });
}

function downloadQr(){
    // cant download if the qr image is not rendered
    if (document.getElementById("qr").src === "") {
        alert("Please generate a QR code first.");
        return;
    }

    // get the image path from the image tag
    let filePath = document.getElementById("qr").src;

    // get the qr_name from the image tag
    let qrName = document.getElementById("qr").name;

    var link = document.createElement("a");
    link.href = filePath; // Replace with the file path or URL
    link.download = qrName; // Replace with the desired filename
    link.click();

}

function validateForm() {
    var emailInput = document.getElementById("email");
    var urlInput = document.getElementById("qr-url");
    
    // Regular expression patterns for email and URL validation
    var emailPattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    var urlPattern = /^(https?|ftp|file):\/\/[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]$|^(www\.)?[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]\.[a-zA-Z]{2,}$/;


    if (!emailPattern.test(emailInput.value)) {
        alert("Please enter a valid email address.");
        emailInput.focus();
        return false;
    }

    if (!urlPattern.test(urlInput.value)) {
        alert("Please enter a valid URL.");
        urlInput.focus();
        return false;
    }

    return true;
}