// Global array for all possible days, accessible to all functions
const allDays = [
    "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche"
];

// Reference to all day select elements (will be populated inside initializeFormLogic)
let daySelects = []; // This needs to be 'let' because it's assigned inside initializeFormLogic


// --- Helper functions ---

// This function updates all the dropdowns
function updateDayOptions() {
    console.log("DEBUG: --- updateDayOptions() called ---");
    const selectedDays = new Set();

    // Phase 1: Determine the intended selection for each *visible* select,
    // and populate the `selectedDays` set.
    console.log("DEBUG: Phase 1: Determining current/initial selections.");
    daySelects.forEach(select => {
        if (!select) return;

        const parentDiv = select.closest('.daysinput');
        const isVisible = !parentDiv || parentDiv.style.display !== 'none';

        let intendedValue = '';
        if (isVisible) {
            intendedValue = select.value || select.dataset.initialValue || '';

            if (intendedValue) {
                selectedDays.add(intendedValue);
                console.log(`DEBUG:   Select '${select.name}' visible, intended: '${intendedValue}'. Added to selectedDays.`);
            } else {
                console.log(`DEBUG:   Select '${select.name}' visible, no intended value.`);
            }
        } else {
            select.value = '';
            select.removeAttribute('data-initial-value');
            console.log(`DEBUG:   Select '${select.name}' is hidden. Value and initial data cleared.`);
        }
    });
    console.log("DEBUG: Phase 1 complete. Current selectedDays set:", Array.from(selectedDays));


    // Phase 2: Populate options and set selections for *all* dropdowns
    console.log("DEBUG: Phase 2: Populating options and setting values.");
    daySelects.forEach(currentSelect => {
        if (!currentSelect) return;

        const parentDiv = currentSelect.closest('.daysinput');
        const isVisible = !parentDiv || parentDiv.style.display !== 'none';

        let targetValue = currentSelect.value || currentSelect.dataset.initialValue || '';

        currentSelect.innerHTML = '';
        console.log(`DEBUG:   Select '${currentSelect.name}' innerHTML cleared.`);

        if (isVisible) {
            allDays.forEach(day => {
                const option = document.createElement('option');
                option.value = day;
                option.textContent = day;

                if (selectedDays.has(day) && day !== targetValue) {
                    option.disabled = true;
                    console.log(`DEBUG:     Option '${day}' for '${currentSelect.name}' disabled.`);
                }
                currentSelect.appendChild(option);
            });
            console.log(`DEBUG:   All options re-created for '${currentSelect.name}'. Total options: ${currentSelect.options.length}.`);

            if (targetValue && Array.from(currentSelect.options).some(opt => opt.value === targetValue && !opt.disabled)) {
                currentSelect.value = targetValue;
                console.log(`DEBUG:   '${currentSelect.name}' value set to its target: '${targetValue}'.`);
            } else {
                const firstAvailableOption = Array.from(currentSelect.options).find(opt => !opt.disabled);
                if (firstAvailableOption) {
                    currentSelect.value = firstAvailableOption.value;
                    console.log(`DEBUG:   '${currentSelect.name}' target unavailable. Set to first available: '${firstAvailableOption.value}'.`);
                } else {
                    currentSelect.value = '';
                    console.log(`DEBUG:   '${currentSelect.name}': No available options, value cleared.`);
                }
            }
            currentSelect.removeAttribute('data-initial-value');
        } else {
            currentSelect.value = '';
            console.log(`DEBUG:   Select '${currentSelect.name}' is hidden, ensuring value is empty.`);
        }
    });
    console.log("DEBUG: --- updateDayOptions() finished ---");
}

function showDayInputs() {
    console.log("DEBUG: showDayInputs() called.");
    const days = parseInt(document.getElementById('days').value, 10);
    const dayDivs = [
        document.getElementById('second'),
        document.getElementById('third'),
        document.getElementById('fourth'),
        document.getElementById('fifth'),
        document.getElementById('sixth'),
        document.getElementById('seventh')
    ];

    for (let i = 0; i < dayDivs.length; i++) {
        const currentDiv = dayDivs[i];
        if (!currentDiv) {
            console.warn(`WARN: Day div at index ${i} is null. Check HTML IDs.`);
            continue;
        }
        const inputsAndSelects = currentDiv.querySelectorAll("input, select");
        if (days && (i + 2 <= days)) {
            currentDiv.style.display = 'block';
            inputsAndSelects.forEach(element => {
                element.required = true;
            });
            console.log(`DEBUG: Day div ${currentDiv.id} set to 'block', inputs required.`);
        } else {
            currentDiv.style.display = 'none';
            inputsAndSelects.forEach(element => {
                element.value = '';
                element.required = false;
                if (element.tagName === 'SELECT') {
                    element.removeAttribute('data-initial-value');
                }
            });
            console.log(`DEBUG: Day div ${currentDiv.id} set to 'none', inputs cleared.`);
        }
    }
    updateDayOptions(); // This call will now work!
}

function Popup() {
    document.getElementById('popup').style.display="block"
}

function Popdown() {
    document.getElementById('popup').style.display="none"
}

function NewTab() {
    const displayedResult = document.getElementById('displayedResult');
    displayedResult.style.display = "none";

    const mainForm = document.getElementById('mainForm');
    if (mainForm) {
        // Create a temporary form element
        const tempForm = document.createElement('form');
        tempForm.method = 'POST';
        tempForm.action = '/result'; // The backend endpoint for new tab results
        tempForm.target = '_blank'; // Crucial for opening in a new tab
        tempForm.style.display = 'none'; // Hide the temporary form

        // Append all input fields from the original form to the temporary form
        // This ensures all form data is transferred
        for (const input of mainForm.elements) {
            if (input.name) { // Only process elements with a 'name' attribute
                const tempInput = document.createElement('input');
                tempInput.type = 'hidden'; // Use hidden inputs for data transfer
                tempInput.name = input.name;
                tempInput.value = input.value;
                tempForm.appendChild(tempInput);
            }
        }

        // Append the temporary form to the body and submit it
        document.body.appendChild(tempForm);
        tempForm.submit();

        // Remove the temporary form after submission
        document.body.removeChild(tempForm);

        // No need to reset mainForm's action/target here, as it was not modified
    } else {
        console.error("ERROR: Main form not found for new tab submission.");
    }
}

function submitForDownload() {
    const form = document.getElementById('mainForm');
    if (form) {
        form.action = '/download';
        form.submit();
        form.action = '/gen';
    } else {
        console.error("ERROR: Main form not found for download submission.");
    }
}


// --- The core initialization function ---
function initializeFormLogic() {
    try {
        console.log("SCRIPT START: initializeFormLogic() triggered.");

        // --- Critical Element Retrieval ---
        const daysInput = document.getElementById('days');
        if (!daysInput) {
            console.error("ERROR: Cannot find #days input. Aborting form setup.");
            return;
        }
        daysInput.addEventListener('input', showDayInputs);
        console.log("DEBUG: #days input found and event listener attached.");

        // Initialize daySelects here once the DOM is ready
        daySelects = [
            document.querySelector('select[name="firstday"]'),
            document.querySelector('select[name="secondday"]'),
            document.querySelector('select[name="thirdday"]'),
            document.querySelector('select[name="fourthday"]'),
            document.querySelector('select[name="fifthday"]'),
            document.querySelector('select[name="sixthday"]'),
            document.querySelector('select[name="seventhday"]')
        ];

        // Verify all select elements are found (added better logging here)
        let allSelectsFound = true;
        for (let i = 0; i < daySelects.length; i++) {
            const expectedName = [
                "firstday", "secondday", "thirdday", "fourthday",
                "fifthday", "sixthday", "seventhday"
            ][i];
            if (!daySelects[i]) {
                console.error(`ERROR: Select element with name="${expectedName}" (index ${i}) not found. Check HTML names.`);
                allSelectsFound = false;
            } else {
                console.log(`DEBUG: Select element found: ${daySelects[i].name}`);
            }
        }
        if (!allSelectsFound) {
            console.error("ERROR: Not all day select elements found. Aborting dropdown population.");
            return;
        }
        console.log("DEBUG: All required input and select elements successfully referenced.");

        // Add change listeners to each select element for dynamic updates
        daySelects.forEach(select => {
            if (select) {
                select.addEventListener('change', updateDayOptions);
                console.log(`DEBUG: Change listener attached to '${select.name}'.`);
            }
        });

        // Initial call to set up form state and populate dropdowns.
        console.log("SCRIPT END: Initial calls to showDayInputs() and updateDayOptions().");
        showDayInputs(); // This will trigger updateDayOptions() for the first time.

    } catch (e) {
        console.error("FATAL ERROR: An uncaught error occurred in initializeFormLogic():", e);
        console.error("Error name:", e.name);
        console.error("Error message:", e.message);
        console.error("Error stack:", e.stack);
        alert("A critical JavaScript error occurred.\n" +
              "Error: " + (e.message || e.name || "Unknown error") + "\n" +
              "Check console (F12) for stack trace.");
    }
}

// --- Universal DOM Ready Handler ---
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initializeFormLogic);
    console.log("DEBUG: Registered DOMContentLoaded listener for initializeFormLogic.");
} else {
    initializeFormLogic();
    console.log("DEBUG: DOM already ready. Calling initializeFormLogic() directly.");
}