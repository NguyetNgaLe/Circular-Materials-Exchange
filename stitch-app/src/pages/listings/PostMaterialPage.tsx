import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'

export default function PostMaterialPage() {
  const nav = useNavigate()
  const store = useStore()
  const [categories, setCategories] = useState<any[]>([])
  const [loading, setLoading] = useState(false)
  const [submitting, setSubmitting] = useState(false)
  const [form, setForm] = useState({
    title: '', categoryId: '', description: '', quantity: '', unit: 'Tấn',
    pricePerUnit: '', location: '', minOrderQuantity: '', packaging: '',
  })
  const set = (k: string, v: string) => setForm(f => ({ ...f, [k]: v }))

  useEffect(() => {
    const fetchCategories = async () => {
      setLoading(true)
      try {
        const data = await store.getCategories()
        setCategories(data)
      } catch (err) {
        console.error('Failed to fetch categories:', err)
      } finally {
        setLoading(false)
      }
    }
    fetchCategories()
  }, [])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setSubmitting(true)
    try {
      const result = await store.createListing({
        ...form,
        quantity: Number(form.quantity),
        pricePerUnit: Number(form.pricePerUnit),
        minOrderQuantity: form.minOrderQuantity ? Number(form.minOrderQuantity) : undefined,
      })
      if (result) {
        alert('Đăng vật liệu thành công!')
        nav('/listings')
      } else {
        alert('Có lỗi xảy ra. Vui lòng thử lại.')
      }
    } catch (err) {
      console.error('Failed to create listing:', err)
      alert('Có lỗi xảy ra. Vui lòng thử lại.')
    } finally {
      setSubmitting(false)
    }
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Đăng vật liệu</p>
          <h1>Đăng Nguồn Cung Mới</h1>
        </div>
      </div>

      <form onSubmit={handleSubmit} className="form-card">
        <div className="form-grid">
          <div className="form-group span-2">
            <label>Tên vật liệu *</label>
            <input required value={form.title} onChange={e => set('title', e.target.value)} placeholder="VD: Nhựa PET tái chế" />
          </div>
          <div className="form-group">
            <label>Danh mục *</label>
            <select required value={form.categoryId} onChange={e => set('categoryId', e.target.value)}>
              <option value="">Chọn danh mục</option>
              {categories.map(c => <option key={c.id} value={c.id}>{c.name}</option>)}
            </select>
          </div>
          <div className="form-group">
            <label>Địa điểm *</label>
            <input required value={form.location} onChange={e => set('location', e.target.value)} placeholder="VD: TP. Hồ Chí Minh" />
          </div>
          <div className="form-group span-2">
            <label>Mô tả *</label>
            <textarea required rows={4} value={form.description} onChange={e => set('description', e.target.value)} placeholder="Mô tả chi tiết về vật liệu..." />
          </div>
          <div className="form-group">
            <label>Số lượng *</label>
            <input required type="number" value={form.quantity} onChange={e => set('quantity', e.target.value)} />
          </div>
          <div className="form-group">
            <label>Đơn vị</label>
            <select value={form.unit} onChange={e => set('unit', e.target.value)}>
              <option>Tấn</option><option>Kg</option><option>Cái</option><option>Mét</option>
            </select>
          </div>
          <div className="form-group">
            <label>Giá (VND) *</label>
            <input required type="number" value={form.pricePerUnit} onChange={e => set('pricePerUnit', e.target.value)} />
          </div>
          <div className="form-group">
            <label>Đơn hàng tối thiểu</label>
            <input type="number" value={form.minOrderQuantity} onChange={e => set('minOrderQuantity', e.target.value)} />
          </div>
          <div className="form-group">
            <label>Đóng gói</label>
            <input value={form.packaging} onChange={e => set('packaging', e.target.value)} placeholder="VD: Bao 25kg" />
          </div>
        </div>
        <div className="form-actions">
          <button type="button" className="btn btn-ghost" onClick={() => nav(-1)}>Hủy</button>
          <button type="submit" className="btn btn-primary" disabled={submitting}>
            {submitting ? 'Đang đăng...' : 'Đăng Vật Liệu'}
          </button>
        </div>
      </form>
    </Layout>
  )
}
