import React, { useState } from 'react';
import { createNewTrack } from '../../api/tracksApi';
import type { NewTrackPayload } from '../../api/tracksApi';

export const NewTrack = () => {
  const [receiverName, setReceiverName] = useState('');
  const [deliveryAddress, setDeliveryAddress] = useState('');
  const [note, setNote] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); 
    setLoading(true);
    setError(null);
    setSuccess(false);

    const token = localStorage.getItem('authToken');
    if (!token) {
      setError('User not authenticated');
      setLoading(false);
      return;
    }

    const payload: NewTrackPayload = {
      receiverName,
      deliveryAddress,
      note,
    };

    try {
      await createNewTrack(payload, token);
      setSuccess(true);
      setReceiverName('');
      setDeliveryAddress('');
      setNote('');
    } catch (err: any) {
      setError(err.message || 'Unexpected error');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="edit">
      <form className="edit__form" onSubmit={handleSubmit}>
        <h2 className="edit__title">CREATE NEW TRACK</h2>
        <div className="edit__form_field">
          <input
            className="edit__form_input"
            placeholder="Receiver"
            value={receiverName}
            onChange={(e) => setReceiverName(e.target.value)}
            required
          />
        </div>
        <div className="edit__form_field">
          <input
            className="edit__form_input"
            placeholder="Address"
            value={deliveryAddress}
            onChange={(e) => setDeliveryAddress(e.target.value)}
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
        {success && <p style={{ color: 'green' }}>Track created successfully!</p>}
      </form>
    </div>
  );
};
