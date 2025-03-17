from flask import Flask, render_template, request, redirect, url_for, jsonify

app = Flask(__name__)

# Define a Book class
class Book:
    def __init__(self, title, author):
        self.title = title
        self.author = author

# In-memory 'database' (a list to store Book objects)
books = [
    Book("The Great Gatsby", "F. Scott Fitzgerald"),
    Book("War of the Worlds", "David Francheski")]

@app.route('/')
def index():
    return render_template('index.html', books=books)

@app.route('/add', methods=['POST'])
def add_book():
    title = request.form.get('title')
    author = request.form.get('author')
    if title and author:
        new_book = Book(title, author)  # Create a new Book object
        books.append(new_book)  # Add it to the list
    return redirect(url_for('index'))

@app.route('/delete/<int:book_id>')
def delete_book(book_id):
    if 0 <= book_id < len(books):
        del books[book_id]  # Deleting the Book object from the list
    return redirect(url_for('index'))

@app.route('/get/<int:book_id>', methods=['GET'])
def get_book(book_id):
    if 0 <= book_id < len(books):
        book = books[book_id]  # Retrieve Book object from the list
        return jsonify({"title": book.title, "author": book.author})
    else:
        return jsonify({"error": "Book not found"}), 404

if __name__ == '__main__':
    app.run(debug=True)
