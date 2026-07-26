import React, { useEffect, useRef, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { Upload, X } from 'lucide-react'
import Layout from '../../components/Layout'
import { useStore } from '../../store'

export default function EditListingPage() {
  const { id } = useParams()
  const nav = useNavigate()
  const store = useStore()
  const fileInputRef = useRef<HTMLInputElement>(null)
  const [categories, setCategories] = useState<any[]>([])
  const [loading, setLoading] = useState(true)
  const [notFound, setNotFound] = useState(false)
  const [submitting, setSubmitting] = useState(false)
  const [uploading, setUploading] = useState(false)
  const [imageUrl, setImageUrl] = useState('')
  const [previewUrl, setPreviewUrl] = useState('')
  const [form, setForm] = useState({
    title: '',
    categoryId: '',
    description: '',
    quantity: '',
    unit: 'Tấn',
    pricePerUnit: '',
    location: '',
    minOrderQuantity: '',
    packaging: '',
  })

  const set = (key: string, value: string) => {
    setForm(current => ({ ...current, [key]: value }))
  }

  useEffect(() => {
    const fetchData = async () => {
      if (!id) {
        setNotFound(true)
        setLoading(false)
        return
      }
      setLoading(true)
      try {
        const [categoryData, listings] = await Promise.all([
          store.getCategories(),
          store.getMyListings(),
        ])
        setCategories(categoryData)
        const listing = listings.find((item: any) => item.id === id)
        if (!listing) {
          setNotFound(true)
          return
        }
        setForm({
          title: listing.title || '',
          categoryId: listing.categoryId || '',
          description: listing.description || '',
          quantity: String(listing.quantity ?? ''),
          unit: listing.unit || 'Tấn',
          pricePerUnit: String(listing.pricePerUnit ?? ''),
          location: listing.location || '',
          minOrderQuantity: listing.minOrderQuantity
            ? String(listing.minOrderQuantity)
            : '',
          packaging: listing.packaging || '',
        })
        const currentImage = listing.imageUrl || ''
        setImageUrl(currentImage)
        setPreviewUrl(currentImage)
      } catch (err) {
        console.error('Failed to load listing:', err)
        setNotFound(true)
      } finally {
        setLoading(false)
      }
    }
    fetchData()
  }, [id])

  const handleImageUpload = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0]
    if (!file) return
    if (file.size > 5 * 1024 * 1024) {
      alert('File quá lớn (tối đa 5MB)')
      return
    }
    if (!['image/jpeg', 'image/png', 'image/gif', 'image/webp'].includes(file.type)) {
      alert('Định dạng không hỗ trợ (chỉ JPG, PNG, GIF, WebP)')
      return
    }

    const reader = new FileReader()
    reader.onload = result => setPreviewUrl(result.target?.result as string)
    reader.readAsDataURL(file)

    setUploading(true)
    try {
      const token = localStorage.getItem('token')
      const formData = new FormData()
      formData.append('file', file)
      const response = await fetch('/api/upload', {
        method: 'POST',
        headers: { Authorization: `Bearer ${token}` },
        body: formData,
      })
      const result = await response.json()
      if (!response.ok || !result.success) {
        throw new Error(result.message || 'Lỗi upload ảnh')
      }
      setImageUrl(result.data.url)
    } catch (err) {
      console.error('Failed to upload image:', err)
      alert('Lỗi upload ảnh')
      setPreviewUrl(imageUrl)
    } finally {
      setUploading(false)
    }
  }

  const removeImage = () => {
    setImageUrl('')
    setPreviewUrl('')
    if (fileInputRef.current) fileInputRef.current.value = ''
  }

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault()
    if (!id) return
    setSubmitting(true)
    try {
      const updated = await store.updateListing(id, {
        title: form.title,
        category_id: form.categoryId,
        description: form.description,
        quantity: Number(form.quantity),
        unit: form.unit,
        price_per_unit: Number(form.pricePerUnit),
        location: form.location,
        min_order_quantity: form.minOrderQuantity
          ? Number(form.minOrderQuantity)
          : 0,
        packaging: form.packaging,
        image_url: imageUrl,
      })
      if (!updated) {
        alert('Không thể cập nhật vật liệu. Vui lòng thử lại.')
        return
      }
      alert('Cập nhật vật liệu thành công!')
      nav('/listings')
    } catch (err) {
      console.error('Failed to update listing:', err)
      alert('Không thể cập nhật vật liệu. Vui lòng thử lại.')
    } finally {
      setSubmitting(false)
    }
  }

  if (loading) {
    return (
      <Layout showSidebar>
        <div className="loading-state">
          <div className="spinner" />
          <p>Đang tải dữ liệu...</p>
        </div>
      </Layout>
    )
  }

  if (notFound) {
    return (
      <Layout showSidebar>
        <div className="empty-state">
          <span className="empty-icon">🔒</span>
          <h3>Không tìm thấy nguồn cung</h3>
          <p>Nguồn cung không tồn tại hoặc bạn không có quyền chỉnh sửa.</p>
          <button className="btn btn-primary" onClick={() => nav('/listings')}>
            Quay lại danh sách
          </button>
        </div>
      </Layout>
    )
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Quản lý vật liệu</p>
          <h1>Chỉnh Sửa Nguồn Cung</h1>
        </div>
      </div>

      <form onSubmit={handleSubmit} className="form-card">
        <div className="form-grid">
          <div className="form-group span-2">
            <label>Tên vật liệu *</label>
            <input required value={form.title} onChange={e => set('title', e.target.value)} />
          </div>
          <div className="form-group">
            <label>Danh mục *</label>
            <select required value={form.categoryId} onChange={e => set('categoryId', e.target.value)}>
              <option value="">Chọn danh mục</option>
              {categories.map(category => (
                <option key={category.id} value={category.id}>{category.name}</option>
              ))}
            </select>
          </div>
          <div className="form-group">
            <label>Địa điểm *</label>
            <input required value={form.location} onChange={e => set('location', e.target.value)} />
          </div>
          <div className="form-group span-2">
            <label>Mô tả *</label>
            <textarea required rows={4} value={form.description} onChange={e => set('description', e.target.value)} />
          </div>
          <div className="form-group">
            <label>Số lượng *</label>
            <input required min="0.000001" step="any" type="number" value={form.quantity} onChange={e => set('quantity', e.target.value)} />
          </div>
          <div className="form-group">
            <label>Đơn vị</label>
            <select value={form.unit} onChange={e => set('unit', e.target.value)}>
              <option>Tấn</option>
              <option>Kg</option>
              <option>Cái</option>
              <option>Mét</option>
            </select>
          </div>
          <div className="form-group">
            <label>Giá (VND) *</label>
            <input required min="0" step="any" type="number" value={form.pricePerUnit} onChange={e => set('pricePerUnit', e.target.value)} />
          </div>
          <div className="form-group">
            <label>Đơn hàng tối thiểu</label>
            <input min="0" step="any" type="number" value={form.minOrderQuantity} onChange={e => set('minOrderQuantity', e.target.value)} />
          </div>
          <div className="form-group">
            <label>Đóng gói</label>
            <input value={form.packaging} onChange={e => set('packaging', e.target.value)} />
          </div>
          <div className="form-group span-2">
            <label>Hình ảnh vật liệu</label>
            <div className="image-upload-area">
              {previewUrl ? (
                <div className="image-preview">
                  <img src={previewUrl} alt="Ảnh vật liệu" />
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
          <button type="button" className="btn btn-ghost" onClick={() => nav('/listings')}>Hủy</button>
          <button type="submit" className="btn btn-primary" disabled={submitting || uploading}>
            {submitting ? 'Đang lưu...' : 'Lưu Thay Đổi'}
          </button>
        </div>
      </form>
    </Layout>
  )
}
