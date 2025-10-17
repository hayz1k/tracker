import React, { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import { fetchSites, deleteSite } from "../../api/sitesApi";
import type { Site } from '../../api/sitesApi';
import { ApiTable } from "../../components/ApiTable/ApiTable";
import './styles.css';
import '../generalStyles.css';

export const ApiPage: React.FC = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 15;

  const [data, setData] = useState<Site[]>([]);
  const [totalPages, setTotalPages] = useState(1);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const [searchTerm, setSearchTerm] = useState('');
  const [appliedFilter, setAppliedFilter] = useState('');

  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem('authToken');
    if (!token) {
      navigate('/login');
      return;
    }

    const loadSites = async () => {
      setLoading(true);
      setError(null);
      try {
        const res = await fetchSites(currentPage, itemsPerPage, token, appliedFilter);
        console.log('Sites API response:', res);
        setData(res.items);
        const totalPagesCalc = res.totalCount > 0 ? res.totalCount : 1;
        setTotalPages(totalPagesCalc);

        if (currentPage > totalPagesCalc) {
          setCurrentPage(totalPagesCalc);
        }
      } catch {
        setError('Ошибка загрузки данных');
      } finally {
        setLoading(false);
      }
    };

    loadSites();
  }, [currentPage, navigate, appliedFilter]);

  const handleDelete = async (id: number) => {
    const token = localStorage.getItem('authToken');
    if (!token) {
      navigate('/login');
      return;
    }

    if (!window.confirm('Вы уверены, что хотите удалить этот сайт?')) return;

    try {
      await deleteSite(id, token);
      setData(prev => prev.filter(site => site.id !== id));
    } catch (err: any) {
      alert(err.message || 'Ошибка при удалении сайта');
    }
  };

  const handlePrev = () => setCurrentPage(p => Math.max(p - 1, 1));
  const handleNext = () => setCurrentPage(p => Math.min(p + 1, totalPages));

  const handleSearchClick = () => {
    setCurrentPage(1);
    setAppliedFilter(searchTerm.trim());
  };

  return (
    <div className="tracks-page section">
      <div className="control-panel">
        <div className="search-fild__block">
          <input
            className="search-fild"
            type="text"
            placeholder="Enter api details"
            value={searchTerm}
            onChange={e => setSearchTerm(e.target.value)}
            onKeyDown={e => {
              if (e.key === 'Enter') {
                handleSearchClick();
              }
            }}
          />
          <button className="search-fild__button" onClick={handleSearchClick}>
            <span className="search-fild__button-dec">G</span>O
          </button>
        </div>
        <div className="control-panel__buttons">
          <Link to="/api/create" className="control-panel__link">
            <p className="control-panel__link_title">NEW API</p>
            <p className="control-panel__link_paragraph">manually</p>
          </Link>
        </div>
      </div>

      {loading && <p>Загрузка...</p>}
      {error && <p style={{ color: 'red' }}>Ошибка: {error}</p>}

      {!loading && !error && (
        <ApiTable
          data={data}
          onEdit={id => navigate(`/api/edit/${id}`)}
          onDelete={handleDelete}
        />
      )}

      <div className="PaginationPanel">
        <div className="GuideButtons">
          <p className="guide-button"><span className="guide-button-dec">E</span> - edit</p>
          <p className="guide-button"><span className="guide-button-dec">U</span> - update</p>
          <p className="guide-button"><span className="guide-button-dec">R</span> - remove</p>
        </div>
        <div className="PaginationButtons">
          <button className="PaginationButton" onClick={handlePrev} disabled={currentPage === 1}>
            ← PREV
          </button>
          <span className="PaginationInfo">
            Page {currentPage} of {totalPages}
          </span>
          <button className="PaginationButton" onClick={handleNext} disabled={currentPage === totalPages}>
            NEXT →
          </button>
        </div>
      </div>
    </div>
  );
};
