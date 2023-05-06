import { memo } from 'react';
import { Gesture, GestureDetector } from 'react-native-gesture-handler';
import Animated, { useAnimatedStyle, useSharedValue, withSpring } from 'react-native-reanimated';
import styled from '@emotion/native';

const AnimatedView = styled(Animated.View)({
  width: 100,
  height: 100,
  borderRadius: 50,
});

const Ball = () => {
  const offset = useSharedValue({ x: 0, y: 0 });
  const isPressed = useSharedValue(false);
  const start = useSharedValue({ x: 0, y: 0 });

  const animatedStyle = useAnimatedStyle(() => {
    return {
      transform: [
        { translateX: offset.value.x },
        { translateY: offset.value.y },
        { scale: withSpring(isPressed.value ? 1.2 : 1) },
      ],
      backgroundColor: isPressed.value ? 'yellow' : 'blue',
    };
  });

  const gesture = Gesture.Pan()
    .onBegin(() => {
      isPressed.value = true;
    })
    .onUpdate((event) => {
      offset.value = {
        x: event.translationX + start.value.x,
        y: event.translationY + start.value.y,
      };
    })
    .onEnd(() => {
      start.value = {
        x: offset.value.x,
        y: offset.value.y,
      };
    })
    .onFinalize(() => {
      isPressed.value = false;
    });

  return (
    <GestureDetector gesture={gesture}>
      <AnimatedView style={animatedStyle} />
    </GestureDetector>
  );
};

export default memo(Ball);
