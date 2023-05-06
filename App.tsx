import { memo } from 'react';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import styled from '@emotion/native';

const StyledGestureHandlerRootView = styled(GestureHandlerRootView)({
  flex: 1,
});

const App = () => {
  return <StyledGestureHandlerRootView></StyledGestureHandlerRootView>;
};

export default memo(App);
