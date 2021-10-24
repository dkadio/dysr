import moment from 'moment';
import {
  Avatar,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  Divider,
  Typography
} from '@mui/material';

const user = {
  avatar: '/static/images/avatars/avatar_6.png',
  city: 'Los Angeles',
  country: 'USA',
  jobTitle: 'Senior Developer',
  name: 'Ele Yinspire me',
  timezone: 'GTM-7'
};

const CodeGeneratorResult = (props) => (
  <Card {...props}>
    <CardContent>
      <Box
        sx={{
          alignItems: 'center',
          display: 'flex',
          flexDirection: 'column'
        }}
      >
        This will be the result of the Code Generator
      </Box>
    </CardContent>
    <Divider />
    <CardActions>
      {/*  <Button color="primary" fullWidth variant="text">
        Upload picture
      </Button>
 */}{' '}
    </CardActions>
  </Card>
);

export default CodeGeneratorResult;
