import axios from 'axios';
import type { Track } from '../types/types';

const API_URL = '/api/orders';

// Получить список треков с пагинацией и фильтром
export const fetchTracks = async (
    page: number,
    pageSize: number,
    token: string,
    filter: string = ''
): Promise<{ tracks: Track[]; total: number }> => {
  const response = await axios.get(API_URL, {
    params: {
      page,       // вместо PageIndex
      limit: pageSize, // вместо PageSize
      filter,     // вместо Filter
    },
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  console.log('API response:', response.data);

  return {
    tracks: response.data.items,
    total: response.data.totalCount,
  };
};


// Удаление трека
export const deleteTrack = async (id: number, token: string): Promise<void> => {
  try {
    await axios.delete(`${API_URL}/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  } catch (error: any) {
    const message =
      error.response?.data?.detail ||
      error.message ||
      'Ошибка при удалении трека';
    throw new Error(message);
  }
};

// Создание нового трека
export type NewTrackPayload = {
  receiverName: string;
  deliveryAddress: string;
  note?: string;
};

export async function createNewTrack(data: NewTrackPayload, token: string) {
  const response = await fetch(API_URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.detail || 'Failed to create track');
  }

  return await response.json();
}

// Получить один трек по id
export const fetchTrackById = async (id: number, token: string) => {
  const res = await fetch(`${API_URL}/${id}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  if (!res.ok) {
    throw new Error(`Ошибка загрузки трека: ${res.status}`);
  }

  return res.json();
};

// Обновить трек
export const updateTrack = async (
  id: number,
  data: {
    note: string;
    deliveryAddress: string;
    receiverName: string;
    customStatus: string;
    orderNumber: string;
  },
  token: string
) => {
  const body = {
    note: data.note,
    deliveryAddress: data.deliveryAddress,
    receiverName: data.receiverName,
    customStatus: data.customStatus,
    status: 0,
    statusGroupId: 0,
  };

  const res = await fetch(`${API_URL}/${id}`, {
    method: 'PUT',
    headers: {
      Authorization: `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(body),
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => ({}));
    const message = errorData.detail || `Ошибка обновления: ${res.status}`;
    throw new Error(message);
  }

  return res.json();
};

// Синхронизация WooCommerce (POST /api/woocommerce/sync)
export const syncWooCommerce = async (token: string) => {
  return axios.post('/api/woocommerce/sync', null, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};
