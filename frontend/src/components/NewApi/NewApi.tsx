import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { createSite } from '../../api/sitesApi';
import './styles.css';

export const NewApi = () => {
  const navigate = useNavigate();

  const [domain, setDomain] = useState('');
  const [consumerKey, setConsumerKey] = useState('');
  const [consumerSecret, setConsumerSecret] = useState('');
  const [note, setNote] = useState('');

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const token = localStorage.getItem('authToken');
    if (!token) {
      navigate('/login');
      return;
    }

    setLoading(true);
    setError(null);

    try {
      await createSite(
        {
          domain,
          consumerKey,
          consumerSecret,
          note,
        },
        token
      );
      navigate('/sites');
    } catch (error) {
      setError('Ошибка создания API');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="edit">
      <form className="edit__form" onSubmit={handleSubmit}>
        <h2 className="edit__title">ADD NEW API</h2>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            placeholder="Consumer Key"
            value={consumerKey}
            onChange={(e) => setConsumerKey(e.target.value)}
            required
          />
        </div>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            placeholder="Consumer Secret"
            value={consumerSecret}
            onChange={(e) => setConsumerSecret(e.target.value)}
            required
          />
        </div>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            placeholder="Domain"
            value={domain}
            onChange={(e) => setDomain(e.target.value)}
            required
          />
        </div>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            placeholder="Custom note"
            value={note}
            onChange={(e) => setNote(e.target.value)}
          />
        </div>

        <button className="edit__form_button" type="submit" disabled={loading}>
          CREAT<span className="edit__form_button-dec">E</span>{loading && '...'}
        </button>

        {error && <p style={{ color: 'red' }}>{error}</p>}
      </form>
    </div>
  );
};
