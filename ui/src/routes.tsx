import { Navigate } from 'react-router-dom';
import DashboardLayout from './components/DashboardLayout';
import MainLayout from './components/MainLayout';
import CodeGeneratorPage from './pages/CodeGeneratorPage';
import CodesList from './pages/CodesList';
import Dashboard from './pages/Dashboard';
import Login from './pages/Login';
import NotFound from './pages/NotFound';
import Register from './pages/Register';
import Settings from './pages/Settings';

const routes = (isLoggedIn) => [
  {
    path: 'app',
    element: <DashboardLayout />,
    children: [
      {
        path: 'generator',
        element: isLoggedIn ? <CodeGeneratorPage /> : <Navigate to="/login" />
      },
      {
        path: 'generator/:id',
        element: isLoggedIn ? <CodeGeneratorPage /> : <Navigate to="/login" />
      },
      { path: 'codes/:id', element: <CodesList /> },
      { path: 'codes', element: <CodesList /> },
      { path: 'dashboard', element: <Dashboard /> },
      { path: 'dashboard/:id', element: <Dashboard /> },
      { path: 'settings', element: <Settings /> },
      { path: '*', element: <Navigate to="/404" /> }
    ]
  },
  {
    path: '/',
    element: <MainLayout />,
    children: [
      { path: 'login', element: <Login /> },
      { path: 'register', element: <Register /> },
      { path: '404', element: <NotFound /> },
      { path: '/', element: <Navigate to="/app/dashboard" /> },
      { path: '*', element: <Navigate to="/404" /> }
    ]
  }
];

export default routes;
