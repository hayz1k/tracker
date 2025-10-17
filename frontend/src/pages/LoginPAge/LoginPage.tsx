import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './styles.css';

interface LoginPageProps {
  onLogin: () => void;
}

export const LoginPage: React.FC<LoginPageProps> = ({ onLogin }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate(); // <-- добавляем хук

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        const text = await response.text();
        setError(text || 'Ошибка авторизации');
        setLoading(false);
        return;
      }

      const data = await response.json();
      localStorage.setItem('authToken', data.token);
      onLogin();

      navigate('/tracks'); // <-- редирект после логина

    } catch {
      setError('Ошибка сети или сервера');
    } finally {
      setLoading(false);
    }
  };

return (
  <div className="login__container">
    <form className="login__form" onSubmit={handleSubmit}>
      <input
        className="login__input"
        placeholder="Email"
        type="email"
        value={email}
        onChange={e => setEmail(e.target.value)}
        required
      />
      <input
        className="login__input"
        placeholder="Password"
        type="password"
        value={password}
        onChange={e => setPassword(e.target.value)}
        required
      />
      <button className="login__button" type="submit" disabled={loading}>
        {loading ? 'Loading...' : 'LOGIN'}
      </button>
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </form>
  </div>
);
};
