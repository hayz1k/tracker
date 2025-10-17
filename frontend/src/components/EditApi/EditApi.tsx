import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { fetchSiteById, updateSite } from '../../api/sitesApi';
import './styles.css';
import "../GeneralEditStyles.css"

export const EditApi: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  const [domain, setDomain] = useState('');
  const [consumerKey, setConsumerKey] = useState('');
  const [consumerSecret, setConsumerSecret] = useState('');
  const [note, setNote] = useState('');

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const loadData = async () => {
      try {
        const token = localStorage.getItem('authToken');
        if (!token) throw new Error('No auth token');

        const data = await fetchSiteById(Number(id), token);

        setDomain(data.domain);
        setConsumerKey(data.consumerKey);
        setConsumerSecret(data.consumerSecret);
        setNote(data.note || '');
      } catch (err: any) {
        setError(err.message || 'Failed to load site data');
      } finally {
        setLoading(false);
      }
    };

    loadData();
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem('authToken');
      if (!token) throw new Error('No auth token');

      await updateSite(
        Number(id),
        { domain, consumerKey, consumerSecret, note },
        token
      );
      navigate('/api');
    } catch (err: any) {
      setError(err.message || 'Failed to update site');
    }
  };

  if (loading) return <p>Загрузка...</p>;
  if (error) return <p style={{ color: 'red' }}>{error}</p>;

  return (
    <div className="edit">
      <form className="edit__form" onSubmit={handleSubmit}>
        <h2 className="edit__title">EDIT API</h2>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            value={domain}
            onChange={(e) => setDomain(e.target.value)}
            placeholder="Domain"
            required
          />
        </div>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            value={consumerKey}
            onChange={(e) => setConsumerKey(e.target.value)}
            placeholder="Consumer Key"
            required
          />
        </div>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            value={consumerSecret}
            onChange={(e) => setConsumerSecret(e.target.value)}
            placeholder="Secret Key"
            required
          />
        </div>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            value={note}
            onChange={(e) => setNote(e.target.value)}
            placeholder="Note"
          />
        </div>

        <button className="edit__form_button" type="submit">
          SAV<span className="edit__form_button-dec">E</span>
        </button>
      </form>
    </div>
  );
};