import axios from "axios";
import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { API_URL } from "../../App";
import toast from "react-hot-toast";

function BookPage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [book, setBook] = useState({
    title: "",
    author: "",
    publishedYear: "",
  });

  useEffect(() => {
    if (id) {
      axios
        .get(`${API_URL}/book/${id}`)
        .then((response) => {
          setBook(response.data);
        })
        .catch((error) => console.error("Error fetching book details:", error));
    }
  }, [id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setBook((prevState) => ({ ...prevState, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (id) {
      // Update existing book
      axios
        .patch(`${API_URL}/book/${id}`, {
          title: book.title,
          author: book.author,
          publishedYear: parseInt(book.publishedYear),
        })
        .then(() => {
          navigate("/");
          toast.success("Updating success..");
        })
        .catch((error) => console.error("Error updating book:", error));
    } else {
      // Add new book
      axios
        .post(`${API_URL}/book/`, {
          title: book.title,
          author: book.author,
          publishedYear: parseInt(book.publishedYear),
        })
        .then(() => {navigate("/");toast.success("Inserting success..");})
        .catch((error) => console.error("Error adding new book:", error));
    }
  };

  return (
    <div className="container">
      <h2 className="text-secondary">{id ? "Edit Book" : "Add Book"}</h2>
      <form onSubmit={handleSubmit} className="mt-4">
        <div className="mb-3">
          <label htmlFor="title" className="form-label">
            Title
          </label>
          <input
            type="text"
            className="form-control"
            id="title"
            name="title"
            value={book.title}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="author" className="form-label">
            Author
          </label>
          <input
            type="text"
            className="form-control"
            id="author"
            name="author"
            value={book.author}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="publishedYear" className="form-label">
            Published Year
          </label>
          <input
            type="number"
            className="form-control"
            id="publishedYear"
            name="publishedYear"
            value={book.publishedYear}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" className="btn btn-primary">
          {id ? "Update" : "Submit"}
        </button>
      </form>
    </div>
  );
}

export default BookPage;
