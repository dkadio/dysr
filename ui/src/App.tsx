import 'react-perfect-scrollbar/dist/css/styles.css';
import Keycloak from 'keycloak-js';
import { ReactKeycloakProvider } from '@react-keycloak/web';
import { ThemeProvider, StyledEngineProvider } from '@mui/material';
import GlobalStyles from './components/GlobalStyles';
import theme from './theme';
import Approuter from './Approuter';
import { RecoilRoot } from 'recoil';

// @ts-ignore
const keycloak = new Keycloak({
  realm: 'dysr',
  url: 'http://localhost:8080/auth',
  clientId: 'dysr'
});

const App = () => (
  <ReactKeycloakProvider authClient={keycloak}>
    <RecoilRoot>
      <StyledEngineProvider injectFirst>
        <ThemeProvider theme={theme}>
          <GlobalStyles />
          <Approuter />
        </ThemeProvider>
      </StyledEngineProvider>
    </RecoilRoot>
  </ReactKeycloakProvider>
);

export default App;
