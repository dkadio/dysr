import { useKeycloak } from '@react-keycloak/web';
import { IconButton } from '@mui/material';
import InputIcon from '@mui/icons-material/Input';

const LoginButton = () => {
  const { keycloak, initialized } = useKeycloak();

  if (keycloak.authenticated) {
    return (
      <IconButton
        color="inherit"
        onClick={() => keycloak.logout()}
        size="large"
      >
        <InputIcon />
      </IconButton>
    );
  }
  return <></>;
};

export default LoginButton;
