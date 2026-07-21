import React, { useState, useEffect } from 'react'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { Plus, Edit, Trash2, Loader2 } from 'lucide-react'

export default function AdminCategoriesPage() {
  const store = useStore()
  const [cats, setCats] = useState<any[]>([])
  const [loading, setLoading] = useState(true)
  const [newName, setNewName] = useState('')
  const [editId, setEditId] = useState<string | null>(null)
  const [editName, setEditName] = useState('')

  useEffect(() => {
    const loadCategories = async () => {
      setLoading(true)
      const data = await store.getCategories()
      setCats(data)
      setLoading(false)
    }
    loadCategories()
  }, [])

  const handleAdd = () => {
    if (newName.trim()) {
      setCats(prev => [...prev, { id: 'cat' + (prev.length + 1), name: newName, icon: 'category' }])
      setNewName('')
    }
  }

  const handleEdit = (id: string, name: string) => {
    setEditId(id)
    setEditName(name)
  }

  const handleSave = () => {
    if (editId && editName.trim()) {
      setCats(prev => prev.map(c => c.id === editId ? { ...c, name: editName } : c))
      setEditId(null)
      setEditName('')
    }
  }

  const handleDelete = (id: string) => {
    if (confirm('Xóa danh mục này?')) {
      setCats(prev => prev.filter(c => c.id !== id))
    }
  }

  if (loading) {
    return (
      <Layout showSidebar>
        <div className="empty-state">
          <Loader2 className="animate-spin" size={32} />
          <p>Đang tải dữ liệu...</p>
        </div>
      </Layout>
    )
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Quản trị</p>
          <h1>Quản Lý Danh Mục Vật Liệu</h1>
        </div>
      </div>

      <div className="panel">
        <div className="form-row" style={{ marginBottom: 16 }}>
          <input placeholder="Tên danh mục mới..." value={newName} onChange={e => setNewName(e.target.value)} style={{ flex: 1 }} />
          <button className="btn btn-primary" onClick={handleAdd}><Plus size={16} /> Thêm</button>
        </div>

        <div className="table-responsive">
          <table className="table">
            <thead>
              <tr><th>ID</th><th>Tên danh mục</th><th>Thao tác</th></tr>
            </thead>
            <tbody>
              {cats.map(c => (
                <tr key={c.id}>
                  <td>{c.id}</td>
                  <td>
                    {editId === c.id ? (
                      <input value={editName} onChange={e => setEditName(e.target.value)} onBlur={handleSave} onKeyDown={e => e.key === 'Enter' && handleSave()} autoFocus />
                    ) : (
                      c.name
                    )}
                  </td>
                  <td>
                    <div className="action-btns">
                      <button className="icon-btn-sm" onClick={() => handleEdit(c.id, c.name)}><Edit size={16} /></button>
                      <button className="icon-btn-sm danger" onClick={() => handleDelete(c.id)}><Trash2 size={16} /></button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </Layout>
  )
}
