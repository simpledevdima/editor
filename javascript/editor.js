document.addEventListener('DOMContentLoaded', function() {
    const editors1 = document.getElementsByClassName("editor single-line-text")
    const editors2 = document.getElementsByClassName("editor checkbox")
    const saveAPIURL = "/editor/api/save"
    // saveData send ajax request with data to api
    function saveData(datatype, i) {
        if (datatype === "checkbox" || editors1[i].innerHTML !== editors1[i].getAttribute("data-default")) {
            let request
            switch (datatype) {
                case "single-line-text":
                    request = JSON.stringify({
                        "conn-id": parseInt(editors1[i].getAttribute("data-conn-id")),
                        "key": editors1[i].getAttribute("data-key"),
                        "value": editors1[i].innerHTML,
                    })
                    break
                case "checkbox":
                    request = JSON.stringify({
                        "conn-id": parseInt(editors2[i].getAttribute("data-conn-id")),
                        "key": editors2[i].getAttribute("data-key"),
                    })
                    break
            }
            let xhr = new XMLHttpRequest()
            xhr.open('POST', saveAPIURL, true)
            xhr.setRequestHeader('Content-type', 'application/json; charset=UTF-8')
            xhr.send(request);
            xhr.onload = function () {
                if(xhr.status === 200) {
                    const response = JSON.parse(xhr.response)
                    if (response.changed) {
                        switch (datatype) {
                            case "single-line-text":
                                editors1[i].setAttribute("data-default", editors1[i].innerHTML)
                                break
                            case "checkbox":
                                editors2[i].setAttribute("data-value", editors2[i].getAttribute("data-value") !== "true")
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
    }
    // single-line-text handlers
    for (let i = 0; i < editors1.length; i++) {
        // send ajax request on focusout editor element
        editors1[i].addEventListener('focusout', () => {
            saveData("single-line-text", i)
        });
        // paste from clipboard as text/plain
        editors1[i].addEventListener("paste", function(e) {
            // cancel paste
            e.preventDefault();
            // get text representation of clipboard
            const text = (e.originalEvent || e).clipboardData.getData('text/plain');
            // insert text manually
            document.execCommand("insertHTML", false, text);
        });
        // save data by enter keypress
        editors1[i].addEventListener('keypress', (e) => {
            if (e.which === 13) {
                e.preventDefault();
                saveData("single-line-text", i)
            }
        });
    }
    // checkbox handlers
    for (let i = 0; i < editors2.length; i++) {
        editors2[i].addEventListener('click', ()=> {
            saveData("checkbox", i)
        })
    }
});
