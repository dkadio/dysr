import { useRoutes } from 'react-router-dom';
import { useKeycloak } from '@react-keycloak/web';
import routes from './routes';
import { OpenAPI } from './gen/api/core/OpenAPI';

const AppRouter = () => {
  const { keycloak, initialized } = useKeycloak();
  const content = useRoutes(routes(keycloak.authenticated));

  if (!keycloak || !initialized) {
    return <>replace with loading page...</>;
  }
  console.log(keycloak);
  OpenAPI.TOKEN = keycloak.idToken;

  return <>{content}</>;
};

export default AppRouter;
