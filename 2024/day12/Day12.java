import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import java.util.Stack;
import java.util.function.BiPredicate;
import java.util.stream.Collectors;

class Day12 {
	public static void main(String[] args) {
		var example = """
				RRRRIICCFF
				RRRRIICCCF
				VVRRRCCFFF
				VVRCCCJFFF
				VVVVCJJCFE
				VVIVCCJJEE
				VVIIICJJEE
				MIIIIIJJEE
				MIIISIJEEE
				MMMISSJEEE""";
		var farm = Farm.parse(example);
		farm.assignPlots2Regions();
		farm.calculateRegionSizes();
		long totalPrice = farm.calculateFencingPrice();
		System.out.println("Total fencing price: " + totalPrice);
		long totalBulkPrice = farm.calculateBulkFencingPrice();
		System.out.println("Total Bulk fencing price: " + totalBulkPrice);

		try {
			String input = Files.readString(Path.of("day12/input.txt"));
			var farm2 = Farm.parse(input);
			farm2.assignPlots2Regions();
			farm2.calculateRegionSizes();
			long totalPrice2 = farm2.calculateFencingPrice();
			System.out.println("Total fencing price: " + totalPrice2);
			// System.out.println("regions : " + farm2.regions.size() + " corners "
			// + farm2.countRegionCorners().size());
			long totalBulkPrice2 = farm2.calculateBulkFencingPrice();
			System.out.println("Total Bulk fencing price: " + totalBulkPrice2);
		} catch (IOException e) {
			System.err.println("Error reading input.txt: " + e.getMessage());
		}
	}

	static record RegionId(long id) {
	}

	public static class Region {
		private long perimeter;
		private long area;

		public Region(long perimeter, long area) {
			this.perimeter = perimeter;
			this.area = area;
		}

		public void incrementPerimeter() {
			this.perimeter++;
		}

		public void incrementArea() {
			this.area++;
		}

		public long perimeter() {
			return perimeter;
		}

		public long area() {
			return area;
		}

		@Override
		public boolean equals(Object o) {
			if (this == o)
				return true;
			if (!(o instanceof Region))
				return false;
			Region region = (Region) o;
			return perimeter == region.perimeter && area == region.area;
		}

		@Override
		public int hashCode() {
			return Objects.hash(perimeter, area);
		}
	}

	static record Coord(int x, int y) {
	}

	public static class Farm {
		private final Map<RegionId, Region> regions;
		private final Map<Coord, RegionId> plots2Region;
		private final char[][] plots;

		public Farm(Map<RegionId, Region> regions, Map<Coord, RegionId> plots2Region, char[][] plots) {
			this.regions = regions;
			this.plots2Region = plots2Region;
			this.plots = plots;
		}

		public static Farm parse(String input) {
			ArrayList<ArrayList<Character>> plots = new ArrayList<>();
			for (String line : input.split("\n")) {
				ArrayList<Character> row = line.chars()
						.mapToObj(ch -> (char) ch)
						.collect(Collectors.toCollection(ArrayList::new));
				if (row.size() > 0) {
					plots.add(row);
				}
			}
			char[][] result = plots.stream()
					.map(row -> row.stream()
							.map(ch -> ch)
							.toArray(Character[]::new))
					.map(row -> {
						char[] primitive = new char[row.length];
						for (int i = 0; i < row.length; i++) {
							primitive[i] = row[i];
						}
						return primitive;
					})
					.toArray(char[][]::new);
			return new Farm(new HashMap<>(), new HashMap<>(), result);
		}

		public void assignPlots2Regions2() {
			var newRegionId = 0;
			for (int y = 0; y < plots.length; y++) {
				for (int x = 0; x < plots[y].length; x++) {
					var left = new Coord(x - 1, y);
					var above = new Coord(x, y - 1);
					var current = new Coord(x, y);
					if (x > 0 && plots[y][x] == plots[y][x - 1]
							&& plots2Region.containsKey(left)) {
						plots2Region.put(current, plots2Region.get(left));
					} else if (y > 0 && plots[y][x] == plots[y - 1][x]
							&& plots2Region.containsKey(above)) {
						plots2Region.put(current, plots2Region.get(above));
					} else {
						plots2Region.put(current, new RegionId(newRegionId));
						newRegionId += 1;
					}
				}
			}
		}

		public void assignPlots2Regions() {
			var newRegionId = 0;
			for (int y = 0; y < plots.length; y++) {
				for (int x = 0; x < plots[y].length; x++) {
					var current = new Coord(x, y);
					// Skip if this plot already has a region
					if (plots2Region.containsKey(current)) {
						continue;
					}

					var left = new Coord(x - 1, y);
					var right = new Coord(x + 1, y);
					var above = new Coord(x, y - 1);
					var below = new Coord(x, y + 1);

					RegionId existingRegion = null;

					// Check each direction for existing regions with same plot type
					if (x > 0 && plots[y][x] == plots[y][x - 1] && plots2Region.containsKey(left)) {
						existingRegion = plots2Region.get(left);
					} else if (x < plots[y].length - 1 && plots[y][x] == plots[y][x + 1]
							&& plots2Region.containsKey(right)) {
						existingRegion = plots2Region.get(right);
					} else if (y > 0 && plots[y][x] == plots[y - 1][x]
							&& plots2Region.containsKey(above)) {
						existingRegion = plots2Region.get(above);
					} else if (y < plots.length - 1 && plots[y][x] == plots[y + 1][x]
							&& plots2Region.containsKey(below)) {
						existingRegion = plots2Region.get(below);
					}

					// If we found an existing region, use it, otherwise create new one
					RegionId regionToUse = existingRegion != null ? existingRegion
							: new RegionId(newRegionId++);

					// Assign the region to current plot
					plots2Region.put(current, regionToUse);

					// Now expand this region to all connected same-type plots
					expandRegion(x, y, regionToUse);
				}
			}
		}

		private void expandRegion(int startX, int startY, RegionId regionId) {
			char plotType = plots[startY][startX];
			Stack<Coord> toCheck = new Stack<>();
			toCheck.push(new Coord(startX, startY));

			while (!toCheck.isEmpty()) {
				Coord current = toCheck.pop();
				int x = current.x();
				int y = current.y();

				// Check all four directions
				checkAndAddToRegion(x - 1, y, plotType, regionId, toCheck); // left
				checkAndAddToRegion(x + 1, y, plotType, regionId, toCheck); // right
				checkAndAddToRegion(x, y - 1, plotType, regionId, toCheck); // above
				checkAndAddToRegion(x, y + 1, plotType, regionId, toCheck); // below
			}
		}

		private void checkAndAddToRegion(int x, int y, char plotType, RegionId regionId, Stack<Coord> toCheck) {
			Coord coord = new Coord(x, y);
			if (isValid(coord) &&
					plots[y][x] == plotType &&
					!plots2Region.containsKey(coord)) {
				plots2Region.put(coord, regionId);
				toCheck.push(coord);
			}
		}

		public void calculateRegionSizes() {
			for (int y = 0; y < plots.length; y++) {
				for (int x = 0; x < plots[y].length; x++) {
					var regionId = plots2Region.get(new Coord(x, y));
					var region = regions.get(regionId);
					if (region == null) {
						region = new Region(0, 0);
						regions.put(regionId, region);
					}

					var left = new Coord(x - 1, y);
					var right = new Coord(x + 1, y);
					var above = new Coord(x, y - 1);
					var below = new Coord(x, y + 1);
					if (!isValid(left) || !hasSameRegion(left, regionId)) {
						region.incrementPerimeter();
					}
					if (!isValid(right) || !hasSameRegion(right, regionId)) {
						region.incrementPerimeter();
					}
					if (!isValid(above) || !hasSameRegion(above, regionId)) {
						region.incrementPerimeter();
					}
					if (!isValid(below) || !hasSameRegion(below, regionId)) {
						region.incrementPerimeter();
					}
					region.incrementArea();
				}
			}
		}

		public long calculateFencingPrice() {
			// System.out.println("Fencing prices by region:");
			return regions.entrySet().stream()
					.peek(entry -> {
						Region region = entry.getValue();
						RegionId id = entry.getKey();
						long price = region.perimeter() * region.area();
						// System.out.printf("Region %d: perimeter=%d, area=%d, price=%d%n",
						// id.id(), region.perimeter(), region.area(), price);
					}).mapToLong(entry -> entry.getValue().perimeter() * entry.getValue().area())
					.sum();

		}

		public long calculateBulkFencingPrice() {
			Map<RegionId, Integer> sideCounts = countRegionCorners();

			return regions.entrySet().stream()
					.mapToLong(entry -> {
						RegionId regionId = entry.getKey();
						Region region = entry.getValue();
						int sides = sideCounts.get(regionId);
						System.out.println(
								"area: %s, sides: %s".formatted(region.area(), sides));
						return (long) sides * region.area();
					})
					.sum();
		}

		public Map<RegionId, Integer> countRegionCorners() {
			Map<RegionId, Integer> cornerCounts = new HashMap<>();

			// Initialize corner count for each region
			for (RegionId regionId : regions.keySet()) {
				cornerCounts.put(regionId, 0);
			}

			// Check each plot in the grid
			for (int y = 0; y < plots.length; y++) {
				for (int x = 0; x < plots[y].length; x++) {
					Coord currentCoord = new Coord(x, y);
					RegionId regionId = plots2Region.get(currentCoord);

					int cornersAtPosition = countCornersAtPosition(currentCoord);
					cornerCounts.merge(regionId, cornersAtPosition, Integer::sum);
				}
			}

			return cornerCounts;
		}

		private int countCornersAtPosition(Coord pos) {
			int corners = 0;
			RegionId currentRegion = plots2Region.get(pos);

			// Define adjacent positions
			Coord above = new Coord(pos.x(), pos.y() - 1);
			Coord below = new Coord(pos.x(), pos.y() + 1);
			Coord left = new Coord(pos.x() - 1, pos.y());
			Coord right = new Coord(pos.x() + 1, pos.y());

			// Define diagonal positions
			Coord topLeft = new Coord(pos.x() - 1, pos.y() - 1);
			Coord topRight = new Coord(pos.x() + 1, pos.y() - 1);
			Coord bottomLeft = new Coord(pos.x() - 1, pos.y() + 1);
			Coord bottomRight = new Coord(pos.x() + 1, pos.y() + 1);

			// Helper function to check if a position has same region
			BiPredicate<Coord, RegionId> hasSameRegion = (coord, region) -> {
				if (!isValid(coord))
					return false;
				return plots2Region.get(coord).equals(region);
			};

			// Helper function to check if a position is different region or invalid
			BiPredicate<Coord, RegionId> isDifferentRegion = (coord, region) -> {
				return !hasSameRegion.test(coord, region);
			};

			// Check each corner configuration
			// Top-Left corner
			if (isDifferentRegion.test(above, currentRegion)
					&& isDifferentRegion.test(left, currentRegion)) {
				corners++;
			} else if (hasSameRegion.test(above, currentRegion)
					&& hasSameRegion.test(left, currentRegion) &&
					isDifferentRegion.test(topLeft, currentRegion)) {
				corners++;
			}

			// Top-Right corner
			if (isDifferentRegion.test(above, currentRegion)
					&& isDifferentRegion.test(right, currentRegion)) {
				corners++;
			} else if (hasSameRegion.test(above, currentRegion)
					&& hasSameRegion.test(right, currentRegion) &&
					isDifferentRegion.test(topRight, currentRegion)) {
				corners++;
			}

			// Bottom-Left corner
			if (isDifferentRegion.test(below, currentRegion)
					&& isDifferentRegion.test(left, currentRegion)) {
				corners++;
			} else if (hasSameRegion.test(below, currentRegion)
					&& hasSameRegion.test(left, currentRegion) &&
					isDifferentRegion.test(bottomLeft, currentRegion)) {
				corners++;
			}

			// Bottom-Right corner
			if (isDifferentRegion.test(below, currentRegion)
					&& isDifferentRegion.test(right, currentRegion)) {
				corners++;
			} else if (hasSameRegion.test(below, currentRegion)
					&& hasSameRegion.test(right, currentRegion) &&
					isDifferentRegion.test(bottomRight, currentRegion)) {
				corners++;
			}

			return corners;
		}

		private boolean hasSameRegion(Coord coord, RegionId regionId) {
			return plots2Region.get(coord).equals(regionId);
		}

		private boolean isValid(Coord c) {
			var xLen = plots[0].length;
			var yLen = plots.length;
			return c.x >= 0 && c.x < xLen && c.y >= 0 && c.y < yLen;
		}

		// Getters if needed
		public Map<RegionId, Region> getRegions() {
			return Collections.unmodifiableMap(regions);
		}

		public Map<Coord, RegionId> getPlots2Region() {
			return Collections.unmodifiableMap(plots2Region);
		}

		public char[][] getPlots() {
			return plots.clone();
		}
	}
}
