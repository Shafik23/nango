document.addEventListener('DOMContentLoaded', function() {
    const apiButton = document.getElementById('apiButton');
    const apiResponse = document.getElementById('apiResponse');

    apiButton.addEventListener('click', async function() {
        try {
            const response = await fetch('/api/hello');
            const data = await response.json();
            
            apiResponse.innerHTML = `
                <strong>Response from API:</strong><br>
                <pre>${JSON.stringify(data, null, 2)}</pre>
            `;
            apiResponse.classList.add('show');
        } catch (error) {
            apiResponse.innerHTML = `<strong>Error:</strong> ${error.message}`;
            apiResponse.classList.add('show');
        }
    });
});