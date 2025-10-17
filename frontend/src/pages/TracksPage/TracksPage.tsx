import { Link, useNavigate } from 'react-router-dom';
import { TracksTable } from '../../components/TracksTable/TracksTable';
import './styles.css';
import '../generalStyles.css';
import type { Track } from '../../types/types';
import { useEffect, useState } from 'react';
import { fetchTracks, deleteTrack, syncWooCommerce } from '../../api/tracksApi';

export const TracksPage = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 10;

  const [tracks, setTracks] = useState<Track[]>([]);
  const [totalPages, setTotalPages] = useState(1);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const [searchTerm, setSearchTerm] = useState('');
  const [appliedFilter, setAppliedFilter] = useState('');

  const navigate = useNavigate();

  // Функция загрузки треков
  const loadTracks = async (page: number, filter: string) => {
    const token = localStorage.getItem('authToken');
    if (!token) {
      navigate('/login');
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const { tracks, total } = await fetchTracks(page, itemsPerPage, token, filter);
      setTracks(tracks);

      const pagesCount = Math.ceil(total / itemsPerPage);
      console.log(total, itemsPerPage)
      setTotalPages(pagesCount);

      if (page > pagesCount) {
        setCurrentPage(pagesCount);
      }
    } catch (err: any) {
      console.error(err);
      setError('Ошибка загрузки данных');
    } finally {
      setLoading(false);
    }
  };

  // Загружаем данные при монтировании и при смене страницы или фильтра
  useEffect(() => {
    loadTracks(currentPage, appliedFilter);
  }, [currentPage, appliedFilter, navigate]);

  const handleDelete = async (id: number) => {
    const token = localStorage.getItem('authToken');
    if (!token) {
      navigate('/login');
      return;
    }

    if (!window.confirm('Вы уверены, что хотите удалить этот трек?')) return;

    try {
      await deleteTrack(id, token);
      setTracks(prev => prev.filter(track => track.id !== id));
    } catch (err: any) {
      alert(err.message || 'Ошибка при удалении трека');
    }
  };

  const handlePrev = () => setCurrentPage(prev => Math.max(prev - 1, 1));
  const handleNext = () => setCurrentPage(prev => Math.min(prev + 1, totalPages));

  const handleSearchClick = () => {
    setCurrentPage(1);
    setAppliedFilter(searchTerm.trim());
  };

  const handleRefresh = async () => {
    const token = localStorage.getItem('authToken');
    if (!token) {
      navigate('/login');
      return;
    }

    setLoading(true);
    setError(null);

    try {
      await syncWooCommerce(token);

      // Подождём чуть-чуть, чтобы данные синхронизировались на сервере
      await new Promise((res) => setTimeout(res, 2000));

      // Перезагружаем первую страницу
      setCurrentPage(1);
      await loadTracks(1, appliedFilter);
    } catch (err: any) {
      console.error(err);
      setError(err.message || 'Ошибка при синхронизации');
    } finally {
      setLoading(false);
    }
  };

  return (
      <div className="tracks-page section">
        <div className="control-panel">
          <div className="search-fild__block">
            <input
                className="search-fild"
                type="text"
                placeholder="Enter track details"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                onKeyDown={(e) => { if (e.key === 'Enter') handleSearchClick(); }}
            />
            <button className="search-fild__button" onClick={handleSearchClick}>
              <span className="search-fild__button-dec">G</span>O
            </button>
          </div>
          <div className="control-panel__buttons">
            <Link to="/tracks/create" className="control-panel__link">
              <p className="control-panel__link_title">NEW TRACK</p>
              <p className="control-panel__link_paragraph">manually</p>
            </Link>
            <button className="control-panel__link" onClick={handleRefresh}>
              <p className="control-panel__link_title">REFRESH</p>
              <p className="control-panel__link_paragraph">update request</p>
            </button>
            <Link to="/tracks/parse" className="control-panel__link">
              <p className="control-panel__link_title">PARSE</p>
              <p className="control-panel__link_paragraph">to the table</p>
            </Link>
          </div>
        </div>

        {loading ? (
            <p>Загрузка...</p>
        ) : error ? (
            <p style={{ color: 'red' }}>{error}</p>
        ) : (
            <TracksTable data={tracks} onDelete={handleDelete} />
        )}

        <div className="PaginationPanel">
          <div className="GuideButtons">
            <p className="guide-button"><span className="guide-button-dec">E</span> - edit</p>
            <p className="guide-button"><span className="guide-button-dec">R</span> - remove</p>
          </div>
          <div className="PaginationButtons">
            <button className="PaginationButton" onClick={handlePrev} disabled={currentPage <= 1}>← PREV</button>
            <span className="PaginationInfo">Page {currentPage} of {totalPages}</span>
            <button className="PaginationButton" onClick={handleNext} disabled={currentPage >= totalPages}>NEXT →</button>
          </div>
        </div>
      </div>
  );
};
