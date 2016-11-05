const INVALID_CLASS = 'form-invalid';

var formFileLabel = document.querySelector('#form-file-label');
var formFileLabelVal = formFileLabel.innerText;
var formFile = document.querySelector('#form-file');
var currentFile = '';

function formFileChange(e) {
    var file = e.target.value.split('\\').pop();
    var fileOk = file !== '';

    currentFile = fileOk ? file : '';
    formFileLabel.innerText = fileOk ? file : formFileLabelVal;

    if (fileOk) {
        formFileLabel.classList.remove(INVALID_CLASS)
    }
}

formFile.addEventListener('change', formFileChange);
formFile.addEventListener('click', formFileChange);

var formTexts = document.querySelectorAll('.form-text');

function formSubmit() {
    var ok = true;

    if (currentFile === '') {
        ok = false;

        formFileLabel.classList.add(INVALID_CLASS);
    }
    else {
        formFileLabel.classList.remove(INVALID_CLASS);
    }

    for (var i = 0; i < formTexts.length; i++) {
        var formText = formTexts[i];

        if (formText.value === '') {
            ok = false;

            formText.classList.add(INVALID_CLASS);
        }
        else {
            formText.classList.remove(INVALID_CLASS);
        }
    }

    return ok;
}

for (var i = 0; i < formTexts.length; i++) {
    var formText = formTexts[i];

    formText.addEventListener('change', function() {
        if (this.value !== '') {
            this.classList.remove(INVALID_CLASS);
        }
    });
}
