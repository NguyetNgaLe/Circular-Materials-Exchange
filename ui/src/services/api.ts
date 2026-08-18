const API_BASE_URL = '/api';

class ApiService {
  private token: string | null = null;

  setToken(token: string) {
    this.token = token;
    localStorage.setItem('token', token);
  }

  getToken(): string | null {
    if (!this.token) {
      this.token = localStorage.getItem('token');
    }
    return this.token;
  }

  clearToken() {
    this.token = null;
    localStorage.removeItem('token');
  }

  private async request(endpoint: string, options: RequestInit = {}) {
    const token = this.getToken();
    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
      ...((options.headers as Record<string, string>) || {}),
    };

    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }

    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
      ...options,
      headers,
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.message || 'Request failed');
    }

    return data;
  }

  // Auth
  async login(email: string, password: string) {
    return this.request('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    });
  }

  async register(name: string, email: string, password: string, phone?: string) {
    return this.request('/auth/register', {
      method: 'POST',
      body: JSON.stringify({ name, email, password, phone }),
    });
  }

  async getMe() {
    return this.request('/auth/me');
  }

  // Categories
  async getCategories() {
    return this.request('/categories');
  }

  // Listings
  async getListings(params?: { category_id?: string; keyword?: string; location?: string }) {
    const query = new URLSearchParams(params as any).toString();
    return this.request(`/listings${query ? `?${query}` : ''}`);
  }

  async getListing(id: string) {
    return this.request(`/listings/${id}`);
  }

  async getMyListings() {
    return this.request('/my/listings');
  }

  async createListing(data: any) {
    return this.request('/listings', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async updateListing(id: string, data: any) {
    return this.request(`/listings/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteListing(id: string) {
    return this.request(`/listings/${id}`, {
      method: 'DELETE',
    });
  }

  // Demands
  async getDemands(params?: { category_id?: string; keyword?: string }) {
    const query = new URLSearchParams(params as any).toString();
    return this.request(`/demands${query ? `?${query}` : ''}`);
  }

  async createDemand(data: any) {
    return this.request('/demands', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  // Offers
  async getOffers(params?: { role?: string; status?: string }) {
    const query = new URLSearchParams(params as any).toString();
    return this.request(`/offers${query ? `?${query}` : ''}`);
  }

  async createOffer(data: any) {
    return this.request('/offers', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async acceptOffer(id: string) {
    return this.request(`/offers/${id}/accept`, {
      method: 'POST',
    });
  }

  async rejectOffer(id: string) {
    return this.request(`/offers/${id}/reject`, {
      method: 'POST',
    });
  }

  // Transactions
  async getTransactions(params?: { status?: string }) {
    const query = new URLSearchParams(params as any).toString();
    return this.request(`/transactions${query ? `?${query}` : ''}`);
  }

  async getTransaction(id: string) {
    return this.request(`/transactions/${id}`);
  }

  async updateTransactionStatus(id: string, status: string, note?: string) {
    return this.request(`/transactions/${id}/status`, {
      method: 'POST',
      body: JSON.stringify({ status, note }),
    });
  }

  // Companies
  async getCompanies(params?: { status?: string }) {
    const query = new URLSearchParams(params as any).toString();
    return this.request(`/companies${query ? `?${query}` : ''}`);
  }

  async getCompany(id: string) {
    return this.request(`/companies/${id}`);
  }

  async createCompany(data: any) {
    return this.request('/companies', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async approveCompany(id: string) {
    return this.request(`/companies/${id}/approve`, {
      method: 'POST',
    });
  }

  async rejectCompany(id: string, reason: string) {
    return this.request(`/companies/${id}/reject`, {
      method: 'POST',
      body: JSON.stringify({ reason }),
    });
  }

  // Reviews
  async getReviews(params?: { reviewee_id?: string }) {
    const query = new URLSearchParams(params as any).toString();
    return this.request(`/reviews${query ? `?${query}` : ''}`);
  }

  async createReview(data: any) {
    return this.request('/reviews', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  // Notifications
  async getNotifications() {
    return this.request('/notifications');
  }

  async markNotificationRead(id: string) {
    return this.request(`/notifications/${id}/read`, {
      method: 'PUT',
    });
  }

  async markAllNotificationsRead() {
    return this.request('/notifications/read-all', {
      method: 'PUT',
    });
  }

  async getUnreadCount() {
    return this.request('/notifications/unread-count');
  }
}

export const api = new ApiService();
