import axios from 'axios';

const API_BASE_URL = '/api';

export interface Site {
  id: number;
  domain: string;
  consumerKey: string;
  consumerSecret: string;
  note: string;
}

export interface SitesResponse {
  pageSize: number;
  pageIndex: number;
  totalCount: number;
  items: Site[];
}

// Получение списка сайтов с фильтром и пагинацией
export const fetchSites = async (
  page: number,
  pageSize: number,
  token: string,
  filter?: string
): Promise<SitesResponse> => {
  console.log('fetchSites called with:', { page, pageSize, filter, token });
  
  const response = await axios.get(`${API_BASE_URL}/sites`, {
    params: {
      Page: page,
      PageSize: pageSize,
      Filter: filter || undefined,
    },
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  console.log('API response data:', response.data);

  return response.data;
};

export const createSite = async (
  site: Omit<Site, 'id'>,
  token: string
): Promise<{ siteId: number }> => {
  const response = await axios.post(
    `${API_BASE_URL}/sites`,
    site,
    {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    }
  );

  return response.data;
};

export const fetchSiteById = async (id: number, token: string) => {
  const res = await axios.get(`${API_BASE_URL}/sites/${id}`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  return res.data;
};

export const updateSite = async (
  id: number,
  body: {
    domain: string;
    consumerKey: string;
    consumerSecret: string;
    note?: string;
  },
  token: string
) => {
  const res = await axios.put(`${API_BASE_URL}/sites/${id}`, body, {
    headers: { Authorization: `Bearer ${token}` },
  });
  return res.data;
};

export const deleteSite = async (id: number, token: string): Promise<void> => {
  await axios.delete(`${API_BASE_URL}/sites/${id}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};
