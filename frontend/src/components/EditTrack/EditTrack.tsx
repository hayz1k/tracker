import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { fetchTrackById, updateTrack } from "../../api/tracksApi";
import type { Track } from "../../types/types";
import "./styles.css";
import "../GeneralEditStyles.css"

export const EditTrack = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  const [track, setTrack] = useState<Track | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [form, setForm] = useState({
    receiverName: "",
    deliveryAddress: "",
    orderNumber: "",
    note: "",
    customStatus: "",
  });

  useEffect(() => {
    const token = localStorage.getItem("authToken");
    if (!token) {
      navigate("/login");
      return;
    }

    const loadTrack = async () => {
      try {
        const data = await fetchTrackById(Number(id), token);
        setTrack(data);
        setForm({
          receiverName: data.receiverName || "",
          deliveryAddress: data.deliveryAddress || "",
          orderNumber: data.orderNumber || "",
          note: data.note || "",
          customStatus: data.customStatus || "",
        });
      } catch (e) {
        setError("Ошибка загрузки данных");
      } finally {
        setLoading(false);
      }
    };

    loadTrack();
  }, [id, navigate]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const token = localStorage.getItem("authToken");
    if (!token) {
      navigate("/login");
      return;
    }

    try {
      await updateTrack(Number(id), form, token);
      alert("Трек успешно обновлен");
      navigate("/tracks");
    } catch (e) {
      alert("Ошибка обновления трека");
    }
  };

  if (loading) return <p>Загрузка...</p>;
  if (error) return <p style={{ color: "red" }}>{error}</p>;

  return (
    <div className="edit">
      <form className="edit__form" onSubmit={handleSubmit}>
        <h2 className="edit__title">EDIT TRACK</h2>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            name="receiverName"
            value={form.receiverName}
            onChange={handleChange}
            placeholder="Pablo Escobar"
          />
        </div>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            name="deliveryAddress"
            value={form.deliveryAddress}
            onChange={handleChange}
            placeholder="London 955124"
          />
        </div>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            name="orderNumber"
            value={form.orderNumber}
            onChange={handleChange}
            placeholder="Order Number"
          />
        </div>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            name="note"
            value={form.note}
            onChange={handleChange}
            placeholder="Custom note"
          />
        </div>

        <div className="edit__form_field">
          <input
            className="edit__form_input"
            name="customStatus"
            value={form.customStatus}
            onChange={handleChange}
            placeholder="Custom status"
          />
        </div>

        {/* <p>{track?.domain || "site.com"}</p> */}

        <button className="edit__form_button" type="submit">
          SAV<span className="edit__form_button-dec">E</span>
        </button>
      </form>
    </div>
  );
};
