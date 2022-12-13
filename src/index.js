import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, createBrowserRouter, RouterProvider } from 'react-router-dom';
import App from './App';
import Keywords from './components/Keywords';
import Keyword from './components/Keyword';
import Login from './components/Login';
import Home from './components/Home.js';
import ErrorPage from './components/ErrorPage';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {index: true, element: <Home /> },
      {
        path: "/keywords",
        element: <Keywords />,
      },
      {
        path: "/keywords/:id",
        element: <Keyword />,
      },
      // {
      //   path: "/admin/movie/0",
      //   element: <EditMovie />,
      // },
      // {
      //   path: "/admin/movie/:id",
      //   element: <EditMovie />,
      // },
      {
        path: "/login",
        element: <Login />,
      },
    ]
  }
])


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
    <App />
      {/* <RouterProvider router={router} /> */}
    </BrowserRouter>
  </React.StrictMode>
);

