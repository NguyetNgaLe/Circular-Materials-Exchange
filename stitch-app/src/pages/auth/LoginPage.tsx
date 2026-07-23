import React, { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import { useStore } from '../../store'
import { api } from '../../services/api'

export default function LoginPage() {
  const store = useStore()
  const nav = useNavigate()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)
  const [tab, setTab] = useState<'login' | 'register'>('login')
  const [regName, setRegName] = useState('')
  const [regEmail, setRegEmail] = useState('')
  const [regPassword, setRegPassword] = useState('')

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError('')

    try {
      const res = await api.login(email, password)
      if (res.success) {
        api.setToken(res.data.token)
        localStorage.setItem('user', JSON.stringify(res.data.user))
        const user = res.data.user
        if (user.role === 'admin') nav('/admin')
        else nav('/dashboard')
      }
    } catch (err: any) {
      setError(err.message || 'Đăng nhập thất bại')
    } finally {
      setLoading(false)
    }
  }

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError('')

    try {
      const res = await api.register(regName, regEmail, regPassword)
      if (res.success) {
        api.setToken(res.data.token)
        localStorage.setItem('user', JSON.stringify(res.data.user))
        nav('/dashboard')
      }
    } catch (err: any) {
      setError(err.message || 'Đăng ký thất bại')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="auth-page">
      <div className="auth-card">
        <div className="auth-header">
          <span className="auth-icon">♻</span>
          <h1>Đăng nhập an toàn</h1>
          <p className="muted">Circular Materials Exchange</p>
        </div>

        <div className="auth-tabs">
          <button className={`auth-tab ${tab === 'login' ? 'active' : ''}`} onClick={() => setTab('login')}>ĐĂNG NHẬP</button>
          <button className={`auth-tab ${tab === 'register' ? 'active' : ''}`} onClick={() => setTab('register')}>ĐĂNG KÝ</button>
        </div>

        {error && <div className="alert alert-error">{error}</div>}

        {tab === 'login' ? (
          <form onSubmit={handleLogin} className="auth-form">
            <div className="form-group">
              <label>Email doanh nghiệp</label>
              <input type="email" value={email} onChange={e => setEmail(e.target.value)} placeholder="admin@congty.vn" />
            </div>
            <div className="form-group">
              <label>Mật khẩu</label>
              <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="••••••••" />
            </div>
            <div className="form-row">
              <label className="checkbox-label"><input type="checkbox" /> Lưu thiết bị này</label>
              <Link to="#" className="link-blue">Quên mật khẩu?</Link>
            </div>
            <button type="submit" className="btn btn-primary btn-full" disabled={loading}>
              {loading ? 'Đang đăng nhập...' : 'ĐĂNG NHẬP'}
            </button>
          </form>
        ) : (
          <form onSubmit={handleRegister} className="auth-form">
            <div className="form-group">
              <label>Tên doanh nghiệp</label>
              <input type="text" value={regName} onChange={e => setRegName(e.target.value)} placeholder="Công ty TNHH ABC" />
            </div>
            <div className="form-group">
              <label>Email quản trị</label>
              <input type="email" value={regEmail} onChange={e => setRegEmail(e.target.value)} placeholder="admin@abc.vn" />
            </div>
            <div className="form-group">
              <label>Mật khẩu</label>
              <input type="password" value={regPassword} onChange={e => setRegPassword(e.target.value)} placeholder="••••••••" />
            </div>
            <button type="submit" className="btn btn-outline btn-full" disabled={loading}>
              {loading ? 'Đang tạo...' : 'TẠO TÀI KHOẢN MỚI'}
            </button>
            <p className="auth-footer muted">Bằng việc đăng ký, bạn đồng ý với <a href="#" className="link-blue">Điều khoản dịch vụ</a> và <a href="#" className="link-blue">Chính sách bảo mật</a>.</p>
          </form>
        )}
      </div>
    </div>
  )
}
