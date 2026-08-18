import React, { useState, useEffect, useRef } from 'react'
import { useNavigate } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { Upload, X } from 'lucide-react'

export default function PostMaterialPage() {
  const nav = useNavigate()
  const store = useStore()
  const fileInputRef = useRef<HTMLInputElement>(null)
  const [categories, setCategories] = useState<any[]>([])
  const [loading, setLoading] = useState(false)
  const [submitting, setSubmitting] = useState(false)
  const [imageUrl, setImageUrl] = useState('')
  const [uploading, setUploading] = useState(false)
  const [previewUrl, setPreviewUrl] = useState('')
  const [company, setCompany] = useState<any>(null)
  const [form, setForm] = useState({
    title: '', categoryId: '', description: '', quantity: '', unit: 'Tấn',
    pricePerUnit: '', location: '', minOrderQuantity: '', packaging: '',
  })
  const set = (k: string, v: string) => setForm(f => ({ ...f, [k]: v }))

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true)
      try {
        const [categoryData] = await Promise.all([store.getCategories()])
        setCategories(categoryData)
        if (store.currentUser?.companyId) {
          const companyData = await store.getCompany(store.currentUser.companyId)
          setCompany(companyData)
        }
      } catch (err) {
        console.error('Failed to fetch data:', err)
      } finally {
        setLoading(false)
      }
    }
    fetchData()
  }, [])

  if (!store.currentUser?.companyId || !company || company.status !== 'verified') {
    return (
      <Layout showSidebar>
        <div className="empty-state">
          <span className="empty-icon">🏢</span>
          <h3>Doanh nghiep chua duoc duyet</h3>
          <p>Ban can co ho do doanh nghiep da duoc admin duyet moi co the dang nguon cung.</p>
          <button className="btn btn-primary" onClick={() => nav('/company')}>Xem ho do doanh nghiep</button>
        </div>
      </Layout>
    )
  }

  const handleImageUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (!file) return

    if (file.size > 5 * 1024 * 1024) {
      alert('File qua lon (toi da 5MB)')
      return
    }
    if (!['image/jpeg', 'image/png', 'image/gif', 'image/webp'].includes(file.type)) {
      alert('Dinh dang khong ho tro (chi jpg, png, gif, webp)')
      return
    }

    const reader = new FileReader()
    reader.onload = event => {
      setPreviewUrl(event.target?.result as string)
    }
    reader.readAsDataURL(file)

    setUploading(true)
    try {
      const token = localStorage.getItem('token')
      const formData = new FormData()
      formData.append('file', file)
      const result = await (await fetch('/api/upload', {
        method: 'POST',
        headers: { 'Authorization': `Bearer ${token}` },
        body: formData,
      })).json()

      if (result.success) {
        setImageUrl(result.data.url)
      } else {
        alert(result.message || 'Loi upload anh')
        setPreviewUrl('')
      }
    } catch {
      alert('Loi upload anh')
      setPreviewUrl('')
    }
    setUploading(false)
  }

  const removeImage = () => {
    setImageUrl('')
    setPreviewUrl('')
    if (fileInputRef.current) fileInputRef.current.value = ''
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setSubmitting(true)
    try {
      const result = await store.createListing({
        title: form.title,
        category_id: form.categoryId,
        description: form.description,
        quantity: Number(form.quantity),
        unit: form.unit,
        price_per_unit: Number(form.pricePerUnit),
        location: form.location,
        min_order_quantity: form.minOrderQuantity ? Number(form.minOrderQuantity) : undefined,
        packaging: form.packaging,
        image_url: imageUrl || undefined,
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
          <div className="form-group span-2">
            <label>Hình ảnh vật liệu</label>
            <div className="image-upload-area">
              {previewUrl ? (
                <div className="image-preview">
                  <img src={previewUrl} alt="Preview" />
                  <button type="button" className="remove-image" onClick={removeImage}>
                    <X size={16} />
                  </button>
                </div>
              ) : (
                <div
                  className="upload-placeholder"
                  onClick={() => fileInputRef.current?.click()}
                  style={{ cursor: uploading ? 'wait' : 'pointer' }}
                >
                  <Upload size={32} />
                  <p>{uploading ? 'Đang tải lên...' : 'Click để chọn ảnh'}</p>
                  <span>JPG, PNG, GIF, WebP (tối đa 5MB)</span>
                </div>
              )}
              <input
                ref={fileInputRef}
                type="file"
                accept="image/jpeg,image/png,image/gif,image/webp"
                onChange={handleImageUpload}
                style={{ display: 'none' }}
              />
            </div>
          </div>
        </div>
        <div className="form-actions">
          <button type="button" className="btn btn-ghost" onClick={() => nav(-1)}>Hủy</button>
          <button type="submit" className="btn btn-primary" disabled={submitting || uploading}>
            {submitting ? 'Đang đăng...' : 'Đăng Vật Liệu'}
          </button>
        </div>
      </form>
    </Layout>
  )
}
