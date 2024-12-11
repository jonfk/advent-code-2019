//JAVA_OPTIONS -Xmx4g

import java.util.ArrayList;
import java.util.HashMap;
import java.util.ListIterator;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import java.util.stream.Stream;

class Day11 {
	public static void main(String[] args) {
		// var example = "1 2024 1 0 9 9 2021976";
		// var input = Input.parse(example);
		// input.blink();
		// System.out.println("input: %s\n".formatted(input));

		var example = "125 17";
		var input = Input.parse(example);
		input.blink(6);

		if (input.stones().size() != 22) {
			System.out.println("Expected = 22, got = %s\n".formatted(input.stones().size()));
		}

		var inputTxt = "0 89741 316108 7641 756 9 7832357 91";
		var inputReal = Input.parse(inputTxt);
		var stones = inputReal.blinkTimes(25);
		System.out.println("You got %s stones after blinking 25 times".formatted(stones));
		var stones2 = inputReal.blinkTimes(75);
		System.out.println("You got %s stones after blinking 75 times".formatted(stones2));
	}

	record BlinkInput(int times, long stone) {
	}

	record Input(ArrayList<Long> stones, HashMap<BlinkInput, Long> cache) {
		public static Input parse(String input) {
			var stones = new ArrayList<Long>();
			for (String numStr : input.split(" ")) {
				var num = Long.parseLong(numStr);
				stones.add(num);
			}
			return new Input(stones, new HashMap<>());
		}

		public long blinkStone(long stone, int times) {
			var input = new BlinkInput(times, stone);
			if (cache().containsKey(input)) {
				return cache().get(input);
			} else if (times == 1) {
				if (stone == 0) {
					return 1;
				}
				var stoneStr = Long.valueOf(stone).toString();
				if (stoneStr.length() % 2 == 0) {
					cache().put(input, 2l);
					return 2;
				}
				return 1;
			} else {
				if (stone == 0) {
					var newRes = blinkStone(1, times - 1);
					cache().put(input, newRes);
					return newRes;
				}
				var stoneStr = Long.valueOf(stone).toString();
				if (stoneStr.length() % 2 == 0) {
					var num1 = Long.parseLong(
							stoneStr.substring(0, stoneStr.length() / 2));
					var num2 = Long.parseLong(
							stoneStr.substring(stoneStr.length() / 2));
					var newRes = blinkStone(num1, times - 1) + blinkStone(num2, times - 1);
					cache().put(input, newRes);
					return newRes;
				}
				var newRes = blinkStone(stone * 2024, times - 1);
				cache().put(input, newRes);
				return newRes;
			}
		}

		public long blinkTimes(int times) {
			return this.stones().stream().mapToLong(stone -> blinkStone(stone, times)).sum();
		}

		public void blink() {
			var newStones = this.stones().stream()
					.parallel()
					.flatMap(stone -> {
						if (stone == 0) {
							return Stream.of(1l);
						}
						var stoneStr = stone.toString();
						if (stoneStr.length() % 2 == 0) {
							var num1 = Long.parseLong(
									stoneStr.substring(0, stoneStr.length() / 2));
							var num2 = Long.parseLong(
									stoneStr.substring(stoneStr.length() / 2));
							return Stream.of(num1, num2);
						}
						return Stream.of(stone * 2024l);
					})
					.collect(Collectors.toCollection(ArrayList::new));
			this.stones().clear();
			this.stones().addAll(newStones);
		}

		public void blink(int times) {

			IntStream.range(0, times).forEach(i -> {
				long startTime = System.nanoTime();
				this.blink();
				long endTime = System.nanoTime();
				long durationMs = (endTime - startTime) / 1_000_000;
				System.out.println(
						"Total execution time for %s th iteration: %s ms".formatted(i,
								durationMs));
			});
		}
	}
}
