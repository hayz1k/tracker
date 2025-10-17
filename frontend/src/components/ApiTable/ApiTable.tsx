import React from 'react';
import {
  useReactTable,
  getCoreRowModel,
  flexRender,
} from '@tanstack/react-table';
import type { ColumnDef } from '@tanstack/react-table';
import type { Site } from '../../api/sitesApi';
import './styles.css';
import '../GenetalTableStyles.css';

type Props = {
  data: Site[];
  onEdit: (id: number) => void;
  onDelete: (id: number) => void;
  onUpdate: (id: number) => void;
};

export const ApiTable: React.FC<Props> = ({ data, onDelete, onEdit, onUpdate }) => {
  const columns: ColumnDef<Site>[] = [
    {
      accessorKey: 'domain',
      header: 'Domain',
    },
    {
      accessorKey: 'consumerKey',
      header: 'Key',
    },
    {
      accessorKey: 'consumerSecret',
      header: 'Secret',
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
            onClick={() => onEdit(row.original.id)}
          >
            E
          </button>
           <button
            className="action-button update-button"
            onClick={() => onUpdate(row.original.id)}
          >
            U
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
