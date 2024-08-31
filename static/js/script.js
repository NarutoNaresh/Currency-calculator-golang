document.addEventListener('DOMContentLoaded', function () {
    const form = document.querySelector('form');

    // Example: Prevent form submission if amount is not provided
    form.addEventListener('submit', function (event) {
        const amountInput = document.getElementById('amount');
        const currencySelect = document.getElementById('currency');
        
        if (amountInput.value.trim() === '') {
            alert('Please enter an amount.');
            event.preventDefault();
            return;
        }

        if (currencySelect.value === '') {
            alert('Please select a currency.');
            event.preventDefault();
            return;
        }
    });
});
