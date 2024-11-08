import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { API_URL } from '../../App';
import { useNavigate } from 'react-router-dom';
import toast from 'react-hot-toast';

function HomePage() {
  const [books, setBooks] = useState([]);
  const navigate = useNavigate()
  const fetchBooks = async () => {
    try {
      const response = await axios.get(API_URL + "/book/");
      setBooks(response.data);
    } catch (error) {
      console.error("Error fetching data", error);
    }
  };
  useEffect(() => {
    fetchBooks()
  }, []);
 
  const handleDelete = async (id) => {
    try {
      if (confirm(`Are you sure you want to delete ${id} ?`)) {
        const response = await axios.delete(`${API_URL}/book/${id}`);
        fetchBooks();
        toast.success(response.data.msg);
      }
    } catch (error) {
      console.error("Error deleing data", error);
    }
  };
  const handleEdit = (id) => {
    if (confirm(`Are you sure you want to update ${id} ?`)){
      navigate(`/book/edit/${id}`);
    }
    
  };
  return (
    <div className="container mt-4">
      <h1 className="mb-4">Books List</h1>
      <div className="row">
        {books.map(book => (
          <div key={book._id} className="col-md-4 mb-4">
            <div className="card">
              <div className="card-body d-flex justify-content-between align-items-center">
                <div>
                  <h5 className="card-title">{book.title}</h5>
                  <h6 className="card-subtitle mb-2 text-muted">{book.author}</h6>
                  <div className="card-text">
                    <p>Published Year: {book.publishedYear}</p>
                  </div>
                </div>
                <div>
                  <button className="btn btn-primary" onClick={() => handleEdit(book._id)}>
                    <i className="bi bi-pencil-square"></i>
                  </button>
                  <button className="btn btn-danger ms-2" onClick={() => handleDelete(book._id)}>
                    <i className="bi bi-trash"></i>
                  </button>
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default HomePage;
