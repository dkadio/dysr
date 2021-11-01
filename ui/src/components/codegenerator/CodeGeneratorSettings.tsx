import { useEffect, useState } from 'react';
import {
  Box,
  Button,
  Card,
  CardContent,
  CardHeader,
  Divider,
  Grid,
  TextField
} from '@mui/material';
import { codeoptions, codevalue, sizeselector } from '../../utils/codestate';
import { useRecoilState } from 'recoil';
import Slider from '@mui/material/Slider';
import ColorPicker from './ColorPicker';

const CodeGeneratorSettings = ({ savecode, ...props }) => {
  const [options, setOptions] = useRecoilState(codeoptions);
  const [value, setValue] = useRecoilState(codevalue);
  const [size, setSize] = useRecoilState(sizeselector);

  const handleChange = (event) => {
    console.log('event', event);
    setOptions({
      ...options,
      [event.target.name]: event.target.value
    });
  };

  const handleValueChange = (event) => {
    console.log('New Address Value', event);
    setValue(event.target.value);
  };

  const handleSize = (event) => {
    setSize(Number(event.target.value));
  };

  const handleDark = (event) => {
    console.log(event);
    setOptions({ ...options, colorDark: event.hex });
  };
  const handleLight = (event) => {
    setOptions({ ...options, colorLight: event.hex });
  };

  useEffect(() => {});

  return (
    <form autoComplete="off" noValidate {...props}>
      <Card>
        <CardHeader subheader="The information can be edited" title="Profile" />
        <Divider />
        <CardContent>
          <Grid container spacing={3}>
            <Grid item lg={12} md={12} xs={12}>
              <TextField
                fullWidth
                placeholder="http://dysr.com"
                value={value}
                variant="standard"
                type="url"
                name="value"
                onChange={handleValueChange}
              />
            </Grid>
            <Grid item lg={8} md={12} xs={12}>
              <Slider
                size="small"
                defaultValue={256}
                aria-label="Small"
                valueLabelDisplay="auto"
                name="width"
                onChange={handleSize}
                min={32}
                max={448}
                step={32}
              />
            </Grid>
            <Grid item lg={2} md={2} xs={6}>
              <ColorPicker
                name="Dark"
                handleChange={handleDark}
                color={options.colorDark}
              />
            </Grid>
            <Grid item lg={2} md={2} xs={6}>
              <ColorPicker
                name="Light"
                handleChange={handleLight}
                color={options.colorLight}
              />
            </Grid>
          </Grid>
        </CardContent>
        <Divider />
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'flex-end',
            p: 2
          }}
        >
          <Button onClick={savecode} color="primary" variant="contained">
            Save
          </Button>
        </Box>
      </Card>
    </form>
  );
};

export default CodeGeneratorSettings;
