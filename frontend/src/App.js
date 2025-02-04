import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navigation from './components/Navigation';
import Home from './pages/Home';
import Login from './pages/Login';
import './App.css';

// Temporary Register component until we implement it
const Register = () => (
  <div className="register">
    <h1>Register</h1>
    {/* Registration form will go here */}
  </div>
);

function App() {
  return (
    <Router>
      <div className="App">
        <Navigation />
        <main className="content">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;
