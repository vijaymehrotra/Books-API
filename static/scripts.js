document.addEventListener('DOMContentLoaded', function() {
    // Fetch and display the list of books when the page loads
    fetchBooks();

    // Handle form submission to add a new book
    document.getElementById('add-book-form').addEventListener('submit', function(event) {
        event.preventDefault(); // Prevent the default form submission

        const title = document.getElementById('title').value;
        const author = document.getElementById('author').value;

        fetch('/create_book', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ title: title, author: author })
        })
        .then(response => response.json())
        .then(data => {
            console.log('Server response:', data); // Debug: Check the server response
            const responseMessage = document.getElementById('response-message');
            responseMessage.textContent = data.message || 'Book added successfully';

            // Optionally reload the book list after adding a new book
            fetchBooks();
        })
        .catch(error => {
            console.error('Error adding book:', error);
        });
    });
});

function fetchBooks() {
    fetch('/books')
        .then(response => response.json())
        .then(data => {
            console.log('Fetched books data:', data); // Debug: Check the fetched data
            const bookList = document.getElementById('book-list');
            bookList.innerHTML = ''; // Clear the list

            // Check if the data format matches expectations
            if (data.books && Array.isArray(data.books)) {
                const ul = document.createElement('ul');
                data.books.forEach(book => {
                    const listItem = document.createElement('li');
                    listItem.classList.add('book-item');
                    const title = book.title || 'Unknown Title';
                    const author = book.author || 'Unknown Author';
                    listItem.innerHTML = `${title} by ${author} 
                    <button onclick="deleteBook(${book.id})">Delete</button>`;
                    ul.appendChild(listItem);
                });
                bookList.appendChild(ul);
            } else {
                console.error('Unexpected data format:', data);
                bookList.innerHTML = 'No books available or error fetching data.';
            }
        })
        .catch(error => {
            console.error('Error fetching book list:', error);
        });
}

function deleteBook(id) {
    fetch(`/delete_book/${id}`, {
        method: 'DELETE',
    })
    .then(response => response.json())
    .then(data => {
        console.log('Server response:', data); // Debug: Check the server response
        fetchBooks(); // Refresh the book list after deletion
    })
    .catch(error => {
        console.error('Error deleting book:', error);
    });
}
