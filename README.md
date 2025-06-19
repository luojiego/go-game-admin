# Game Management Framework

A modern, flexible game backend management system built with Go and AdminLTE, featuring dynamic routing, RBAC, and i18n support.

## 🌟 Features

### Core Framework
- **Go Backend**: High-performance server implementation
- **AdminLTE Integration**: Modern and responsive admin interface
- **Template System**: Modular components with hot-reloading support
- **Dynamic Routing**: Add and modify routes without service restart

### Security & Management
- **Role-Based Access Control**: Fine-grained permission management
- **Dynamic Route Management**: Web-based route configuration
- **Audit Logging**: Comprehensive activity tracking

### Developer Tools
- **Page Generator**: CLI tool for rapid page creation
- **Environment-aware Logger**: Configurable logging levels per environment
- **Hot Reload**: Development-friendly setup with live updates

### Internationalization
- Multi-language support
- Easy language pack integration
- Dynamic language switching

### Demo Components
- User Management System
- Transaction/Recharge Tracking
- User Analytics Dashboard
- Customizable Widgets

## 🚀 Quick Start

### Prerequisites
- Go 1.16+
- Node.js 14+
- MySQL 5.7+

### Installation
```bash
# Clone the repository
git clone [repository-url]

# Install dependencies
go mod tidy
npm install

# Configure environment
cp .env.example .env

# Start the server
go run main.go
```

### Creating New Pages
```bash
# Generate a new page
./tools/generate-page.sh --name user-list --template crud

# Generate associated JavaScript
./tools/generate-js.sh --name user-list
```

## 📚 Documentation

### Template System
The framework uses a modular template system with the following components:
- `header.html`: Common header elements
- `nav.html`: Navigation menu
- `scripts.html`: Shared JavaScript resources
- `footer.html`: Common footer elements

### Route Management
Routes can be managed through the admin interface at `/admin/routes`:
- Add new routes
- Enable/disable routes
- Modify route parameters
- No service restart required

### Logging System
```javascript
// JavaScript logging example
logger.configure({
    development: ['debug', 'info', 'warn', 'error'],
    production: ['warn', 'error']
});
```

## 🛠 Development

### Project Structure
```
├── cmd/
│   └── main.go
├── internal/
│   ├── auth/
│   ├── router/
│   └── handler/
├── templates/
│   ├── layout/
│   └── pages/
├── static/
│   └── dist/
└── tools/
    └── generators/
```

### Adding New Features
1. Create template in `templates/pages/`
2. Add JavaScript in `static/dist/js/pages/`
3. Configure route in admin panel
4. Implement backend handler

## 📄 License

MIT License - see the [LICENSE](LICENSE) file for details

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details. 