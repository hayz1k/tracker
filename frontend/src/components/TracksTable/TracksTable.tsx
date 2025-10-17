import React from 'react';
import {
  useReactTable,
  getCoreRowModel,
  flexRender,
} from '@tanstack/react-table';
import type { ColumnDef } from '@tanstack/react-table';
import type { Track } from '../../types/types';
import { useNavigate } from 'react-router-dom';
import './styles.css';
import '../GenetalTableStyles.css';

type Props = {
  data: Track[];
  onDelete: (id: number) => void;
};

export const TracksTable: React.FC<Props> = ({ data, onDelete }) => {
  const navigate = useNavigate();

  const columns: ColumnDef<Track>[] = [
    {
      accessorKey: 'order_id',
      header: 'OrderID',
    },
    {
      accessorKey: 'track_number',
      header: 'Track',
    },
    {
      accessorKey: 'first_name',
      header: 'Name',
    }, {
          accessorKey: 'last_name',
          header: 'Surname',
      },
    {
      header: 'Domain',
      accessorKey: 'site.domain',
    },
    {
      accessorKey: 'note',
      header: 'Note',
    },
    {
      id: 'actions',
      header: 'Actions',
      cell: ({ row }) => (
        <div className="table__actions">
          <button
            className="action-button edit-button"
            onClick={() => navigate(`/tracks/edit/${row.original.id}`)}
          >
            E
          </button>
          <button
            className="action-button remove-button"
            onClick={() => onDelete(row.original.id)}
          >
            R
          </button>
        </div>
      ),
    },
  ];

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <table className="table">
      <thead className="table__titles">
        {table.getHeaderGroups().map((headerGroup) => (
          <tr key={headerGroup.id}>
            {headerGroup.headers.map((header) => (
              <th key={header.id}>
                {header.isPlaceholder
                  ? null
                  : flexRender(header.column.columnDef.header, header.getContext())}
              </th>
            ))}
          </tr>
        ))}
      </thead>
      <tbody className="table__fields">
        {table.getRowModel().rows.map((row) => (
          <tr className="table__el" key={row.id}>
            {row.getVisibleCells().map((cell) => (
              <td key={cell.id}>
                {flexRender(cell.column.columnDef.cell, cell.getContext())}
              </td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
};
