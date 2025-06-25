function showDayInputs() {
    const days = parseInt(document.getElementById('days').value, 10);
    const dayDivs = [
        document.getElementById('second'),
        document.getElementById('third'),
        document.getElementById('fourth'),
        document.getElementById('fifth'),
        document.getElementById('sixth'),
        document.getElementById('seventh')
    ];

    // Loop through each day div
    for (let i = 0; i < dayDivs.length; i++) {
        const currentDiv = dayDivs[i];
        const inputsAndSelects = currentDiv.querySelectorAll("input, select");
        if (days && (i + 2 <= days)) {
            currentDiv.style.display = 'block';
            inputsAndSelects.forEach(element => {
                element.required = true;
            });
        } else {
            currentDiv.style.display = 'none';
            inputsAndSelects.forEach(element => {
                element.value = '';
                element.required = false;
            });
        }
    }
}

function Popup() {
    document.getElementById('popup').style.display="block"
}

function Popdown() {
    document.getElementById('popup').style.display="none"
}

function submitForDownload() {
        const form = document.getElementById('mainForm'); // Get a reference to our form
        form.action = '/download'; // STEP 3a: Change the form's submission URL
        form.submit();             // STEP 3b: Programmatically submit the form
        form.action = '/gen';      // STEP 3c: Reset the form's action back
                                   //          for subsequent "Sur cette page" or "Dans un nouvel onglet" clicks
    }

document.addEventListener('DOMContentLoaded', function() {
    const daysInput = document.getElementById('days');
    if (daysInput) {
        daysInput.addEventListener('input', showDayInputs);
    }
    showDayInputs(); // Call initially to set correct visibility on load


    const daySelects = [
        document.querySelector('select[name="firstday"]'),
        document.querySelector('select[name="secondday"]'),
        document.querySelector('select[name="thirdday"]'),
        document.querySelector('select[name="fourthday"]'),
        document.querySelector('select[name="fifthday"]'),
        document.querySelector('select[name="sixthday"]'),
        document.querySelector('select[name="seventhday"]')
    ];
    const allDays = [
        "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche"
    ];

    // This function updates all the dropdowns
    function updateDayOptions() {
        const selectedDays = new Set(); // Keep track of days already chosen


        // First, figure out which days are currently selected across all dropdowns
        daySelects.forEach(select => {
            if (select && select.value) {
                selectedDays.add(select.value);
            }
        });

        // Now, go through each dropdown and update its options
        daySelects.forEach(currentSelect => {
            if (!currentSelect) return; // Skip if a dropdown element isn't found

            const previouslySelectedValue = currentSelect.value; // Remember what was chosen in this dropdown

            // Clear out old options
            currentSelect.innerHTML = '';

            // Add back all the days, disabling those already selected elsewhere
            allDays.forEach(day => {
                const option = document.createElement('option');
                option.value = day;
                option.textContent = day;

                // If this day is already picked in another dropdown AND it's not the day currently
                // chosen in *this* dropdown, then disable it.
                if (selectedDays.has(day) && day !== previouslySelectedValue) {
                    option.disabled = true;
                }
                currentSelect.appendChild(option);
            });

            // Try to set the dropdown back to its previously selected value
            // Only if that value is still available (not disabled)
            if (previouslySelectedValue && Array.from(currentSelect.options).some(opt => opt.value === previouslySelectedValue && !opt.disabled)) {
                currentSelect.value = previouslySelectedValue;
            } else if (previouslySelectedValue) {
                // If the previous value became disabled, select the first available option
                const availableOption = Array.from(currentSelect.options).find(opt => !opt.disabled);
                if (availableOption) {
                    currentSelect.value = availableOption.value;
                } else {
                    currentSelect.value = ''; // If no options are left, clear it
                }
            } else {
                currentSelect.value = ''; // Ensure no default if nothing was selected
            }
        });
    }

    // Add an event listener to each dropdown
    // So when a user changes a selection, the options in other dropdowns update
    daySelects.forEach(select => {
        if (select) {
            select.addEventListener('change', updateDayOptions);
        }
    });

    // Run the function once when the page loads
    // This sets up the initial state of the dropdowns
    updateDayOptions();
});