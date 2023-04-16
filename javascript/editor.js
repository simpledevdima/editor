document.addEventListener('DOMContentLoaded', function() {
    document.execCommand("defaultParagraphSeparator", false, "p");
    const editors1 = document.getElementsByClassName("editor input-text")
    const editors2 = document.getElementsByClassName("editor checkbox")
    const editors3 = document.getElementsByClassName("editor content-html")
    const editors4 = document.getElementsByClassName("editor textarea")
    const editors5 = document.getElementsByClassName("editor select")
    const saveAPIURL = "/editor/api/save"
    // saveData send ajax request with data to api
    function saveData(datatype, request, i) {
        let xhr = new XMLHttpRequest()
        xhr.open('POST', saveAPIURL, true)
        xhr.setRequestHeader('Content-type', 'application/json; charset=UTF-8')
        xhr.send(request);
        xhr.onload = function () {
            if(xhr.status === 200) {
                const response = JSON.parse(xhr.response)
                if (response.changed) {
                    switch (datatype) {
                        case "input-text":
                            responseProcessingInputType(i)
                            break
                        case "checkbox":
                            responseProcessingCheckBox(i)
                            break
                        case "content-html":
                            responseProcessingContentHTML(i)
                            break
                        case "textarea":
                            responseProcessingTextArea(i)
                            break
                    }
                } else {
                    switch (response.error) {
                        case "wrong_connID":
                            alert("Error: invalid database connection id!")
                            break
                        case "wrong_key":
                            alert("Error: wrong key!")
                            break
                    }
                }
            }
        }
    }
    // response processing input-type
    function responseProcessingInputType(i) {
        editors1[i].setAttribute("data-default", editors1[i].value)
    }
    // save input text
    function saveInputText(i) {
        // check data
        if (editors1[i].value !== editors1[i].getAttribute("data-default")) {
            // construct json object
            const request = JSON.stringify({
                "conn-id": parseInt(editors1[i].getAttribute("data-conn-id")),
                "key": editors1[i].getAttribute("data-key"),
                "value": editors1[i].value,
            })
            // send data
            saveData("input-text", request, i)
        }
    }
    // input-text handlers
    for (let i = 0; i < editors1.length; i++) {
        // send ajax request on focusout editor element
        editors1[i].addEventListener('focusout', () => {
            saveInputText(i)
        });
        // save data by enter keypress
        editors1[i].addEventListener('keypress', (e) => {
            if (e.which === 13) {
                saveInputText(i)
            }
        });
    }
    // response processing checkbox
    function responseProcessingCheckBox(i) {
        editors2[i].setAttribute("data-value", editors2[i].getAttribute("data-value") !== "true")
        if (editors2[i].getAttribute("data-value") === "true") {
            editors2[i].innerHTML = "✔"
        } else {
            editors2[i].innerHTML = ""
        }
    }
    // save checkbox
    function saveCheckBox(i) {
        // construct json object
        const request = JSON.stringify({
            "conn-id": parseInt(editors2[i].getAttribute("data-conn-id")),
            "key": editors2[i].getAttribute("data-key"),
        })
        // send data
        saveData("checkbox", request, i)
    }
    // checkbox handlers
    for (let i = 0; i < editors2.length; i++) {
        editors2[i].addEventListener('click', ()=> {
            saveCheckBox(i)
        })
    }

    // set default-data for content-html
    for (let i = 0; i < editors3.length; i++) {
        editors3[i].setAttribute("data-default", editors3[i].innerHTML)
    }
    // response processing input-type
    function responseProcessingContentHTML(i) {
        editors3[i].setAttribute("data-default", editors3[i].innerHTML)
    }
    // save content html
    function saveContentHTML(i) {
        // check data
        if (editors3[i].innerHTML !== editors3[i].getAttribute("data-default")) {
            // construct json object
            const request = JSON.stringify({
                "conn-id": parseInt(editors3[i].getAttribute("data-conn-id")),
                "key": editors3[i].getAttribute("data-key"),
                "value": editors3[i].innerHTML,
            })
            // send data
            saveData("content-html", request, i)
        }
    }
    // content-html handlers
    for (let i = 0; i < editors3.length; i++) {
        // send ajax request on focusout editor element
        editors3[i].addEventListener('focusout', () => {
            saveContentHTML(i)
        });
    }

    // response processing textarea
    function responseProcessingTextArea(i) {
        editors4[i].setAttribute("data-default", editors4[i].value)
    }
    // save textarea
    function saveTextArea(i) {
        // check data
        if (editors4[i].value !== editors4[i].getAttribute("data-default")) {
            // construct json object
            const request = JSON.stringify({
                "conn-id": parseInt(editors4[i].getAttribute("data-conn-id")),
                "key": editors4[i].getAttribute("data-key"),
                "value": editors4[i].value,
            })
            // send data
            saveData("textarea", request, i)
        }
    }
    // textarea handlers
    for (let i = 0; i < editors4.length; i++) {
        // send ajax request on focusout editor element
        editors4[i].addEventListener('focusout', () => {
            saveTextArea(i)
        });
    }

    // save select
    function saveSelect(i) {
        // construct json object
        const request = JSON.stringify({
            "conn-id": parseInt(editors5[i].getAttribute("data-conn-id")),
            "key": editors5[i].getAttribute("data-key"),
            "value": editors5[i].value,
        })
        // send data
        saveData("select", request, i)
    }
    // select handlers
    for (let i = 0; i < editors5.length; i++) {
        editors5[i].addEventListener('change', ()=> {
            saveSelect(i)
        })
    }


    // // single-line-text handlers
    // for (let i = 0; i < editors1.length; i++) {
    //     // send ajax request on focusout editor element
    //     editors1[i].addEventListener('focusout', () => {
    //         editors1[i].innerHTML = editors1[i].innerHTML.trim()
    //         saveData("single-line-text", i)
    //     });
    //     // paste from clipboard as text/plain
    //     editors1[i].addEventListener("paste", function(e) {
    //         // cancel paste
    //         e.preventDefault();
    //         // get text representation of clipboard
    //         const text = (e.originalEvent || e).clipboardData.getData('text/plain');
    //         // insert text manually
    //         document.execCommand("insertHTML", false, text);
    //     });
    //     // save data by enter keypress
    //     editors1[i].addEventListener('keypress', (e) => {
    //         if (e.which === 13) {
    //             e.preventDefault();
    //             saveData("single-line-text", i)
    //         }
    //     });
    // }
});
