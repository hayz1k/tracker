import React, { useState } from 'react';
import { Routes, Route, Navigate, useLocation } from 'react-router-dom';
import { Header } from './components/Header/Header';
import { Footer } from './components/Footer/Footer';
import { TracksPage } from './pages/TracksPage/TracksPage';
import { ApiPage } from './pages/ApiPage/ApiPage';
import { TGApiPAge } from './pages/TGApiPage/TGApiPage';
import { NewTrack } from './components/NewTrack/NewTrack';
import { EditTrack } from './components/EditTrack/EditTrack';
import { NewApi } from './components/NewApi/NewApi';
import { EditApi } from './components/EditApi/EditApi';
import { LoginPage } from './pages/LoginPAge/LoginPage';
import './styles/styles.css';

const App: React.FC = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(() => !!localStorage.getItem('authToken'));
  const location = useLocation();

  const handleLogout = () => {
    localStorage.removeItem('authToken');
    setIsAuthenticated(false);
  };

  const showHeaderFooter = location.pathname !== '/login';

  return (
    <div className="Page">
      {showHeaderFooter && <Header />}
      <div className="Content">
        <Routes>
          {/* Auth */}
          <Route path="/login" element={<LoginPage onLogin={() => setIsAuthenticated(true)} />} />

          {/* Tracks */}
          <Route
            path="/tracks"
            element={isAuthenticated ? <TracksPage /> : <Navigate to="/login" />}
          />
          <Route
            path="/tracks/create"
            element={isAuthenticated ? <NewTrack /> : <Navigate to="/login" />}
          />
          <Route
            path="/tracks/edit/:id"
            element={isAuthenticated ? <EditTrack /> : <Navigate to="/login" />}
          />

          {/* API */}
          <Route
            path="/api"
            element={isAuthenticated ? <ApiPage /> : <Navigate to="/login" />}
          />
          <Route
            path="/api/create"
            element={isAuthenticated ? <NewApi /> : <Navigate to="/login" />}
          />
          <Route
            path="/api/edit/:id"
            element={isAuthenticated ? <EditApi /> : <Navigate to="/login" />}
          />

          {/* Telegram API */}
          <Route
            path="/tg-api"
            element={isAuthenticated ? <TGApiPAge /> : <Navigate to="/login" />}
          />

          {/* Redirect */}
          <Route
            path="*"
            element={isAuthenticated ? <Navigate to="/tracks" /> : <Navigate to="/login" />}
          />
        </Routes>
      </div>
      {showHeaderFooter && <Footer />}
      {isAuthenticated && showHeaderFooter && (
        <button
  onClick={handleLogout}
  style={{
    position: 'fixed',
    bottom: 10,
    right: 10,
    padding: '8px 16px',
    background: '#c00',
    color: '#fff',
    border: 'none',
    borderRadius: '4px',
    cursor: 'pointer',
    fontFamily: 'var(--main-font)',
    fontWeight: 800,
  }}
>
          Logout
        </button>
      )}
    </div>
  );
};

export default App;
