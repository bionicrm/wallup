const INVALID_CLASS = 'form-invalid';

var formFileLabel = document.querySelector('#form-file-label');
var formFileLabelInitVal = formFileLabel.innerText;
var formFile = document.querySelector('#form-file');
var currentFile = '';

function formFileChange(e) {
    var file = e.target.value.split('\\').pop();
    var fileOk = file !== '';

    currentFile = fileOk ? file : '';
    formFileLabel.innerText = fileOk ? file : formFileLabelInitVal;

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

    var validate = function(formText, int, min, inclusive) {
        if (!validateNumberField(formText, int, min, inclusive)) {
            ok = false;
            formText.classList.add(INVALID_CLASS);
        }
        else {
            formText.classList.remove(INVALID_CLASS);
        }
    };

    validate(document.querySelector('#form-x'), true, 0, true);
    validate(document.querySelector('#form-y'), true, 0, true);
    validate(document.querySelector('#form-width-l'), true, 1, true);
    validate(document.querySelector('#form-width-r'), true, 1, true);
    validate(document.querySelector('#form-height'), true, 1, true);
    validate(document.querySelector('#form-scale'), false, 0, false);
    validate(document.querySelector('#form-gap'), true, 0, true);

    return ok;
}

function validateNumberField(formText, int, min, inclusive) {
    // validate presence
    if (formText.value === '') {
        return false;
    }

    var fVal = parseFloat(formText.value);

    // validate min
    if ((inclusive && fVal < min) || (!inclusive && fVal <= min)) {
        return false;
    }

    // validate int
    return !(int && fVal !== Math.floor(fVal));
}

for (var i = 0; i < formTexts.length; i++) {
    var formText = formTexts[i];

    formText.addEventListener('change', function() {
        if (this.value !== '') {
            this.classList.remove(INVALID_CLASS);
        }
    });
}
